package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type CustomGatewayMuxFunc func(*runtime.ServeMux) (http.Handler, error)

type ServerOptions struct {
	Logger           zerolog.Logger
	GrpcGatewayPort  string
	GrpcInternalPort string
	HTTPGatewayPort  string

	ShowDetailInternalError bool

	GrpcGatewayOption  []grpc.ServerOption
	GrpcInternalOption []grpc.ServerOption
}

type MonoServer interface {
	Run() error

	InternalClientConn() *grpc.ClientConn

	CustomGatewayMux(CustomGatewayMuxFunc)

	RegisterGrpcService(sd *grpc.ServiceDesc, ss interface{})
	RegisterGateway(ctx context.Context)

	CompleteServer() error
}

func NewIncompletedServer(opts ServerOptions) (MonoServer, error) {
	logger := opts.Logger

	logger.Info().Msgf("Create gateway client conn to port %v", opts.GrpcGatewayPort)
	gwClientConn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("0.0.0.0:%v", opts.GrpcGatewayPort),
		// grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	logger.Info().Msgf("Create internal client conn to port %v", opts.GrpcInternalPort)
	intClientConn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("0.0.0.0:%v", opts.GrpcInternalPort),
		// grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &incompletedServer{
		internalGrpcServer: grpc.NewServer(opts.GrpcInternalOption...),
		gatewayGrpcServer:  grpc.NewServer(opts.GrpcGatewayOption...),
		internalClientConn: intClientConn,
		gatewayClientConn:  gwClientConn,
		opts:               opts,
		gatewayMux: runtime.NewServeMux(
			runtime.WithErrorHandler(FormatRestError(opts.ShowDetailInternalError)),
		),
		logger: opts.Logger,
	}, nil
}

type incompletedServer struct {
	opts   ServerOptions
	logger zerolog.Logger

	internalGrpcServer *grpc.Server
	gatewayGrpcServer  *grpc.Server

	internalClientConn *grpc.ClientConn
	gatewayClientConn  *grpc.ClientConn

	gatewayMux         *runtime.ServeMux
	fnCustomGatewayMux CustomGatewayMuxFunc
}

func (s *incompletedServer) RegisterGrpcService(sd *grpc.ServiceDesc, ss interface{}) {
	s.gatewayGrpcServer.RegisterService(sd, ss)
	s.internalGrpcServer.RegisterService(sd, ss)
}

func (s *incompletedServer) InternalGrpcServer() *grpc.Server {
	return s.internalGrpcServer
}

func (s *incompletedServer) GatewayGrpcServer() *grpc.Server {
	return s.gatewayGrpcServer
}

func (s *incompletedServer) GatewayServeMux() *runtime.ServeMux {
	return s.gatewayMux
}

func (s *incompletedServer) InternalClientConn() *grpc.ClientConn {
	return s.internalClientConn
}

func (s *incompletedServer) GatewayClientConn() *grpc.ClientConn {
	return s.gatewayClientConn
}

func (s *incompletedServer) Run() error {
	reflection.Register(s.gatewayGrpcServer)
	reflection.Register(s.internalGrpcServer)
	s.logger.Info().Msgf("Start listen grpc at %v", s.opts.GrpcGatewayPort)
	go func() {
		err := s.listenGrpcService(s.gatewayGrpcServer, context.Background(), s.opts.GrpcGatewayPort)
		if err != nil {
			s.logger.Fatal().Err(err).Msg("")
		}
	}()

	s.logger.Info().Msgf("Start listen internal grpc at %v", s.opts.GrpcGatewayPort)
	go func() {
		err := s.listenGrpcService(s.internalGrpcServer, context.Background(), s.opts.GrpcInternalPort)
		if err != nil {
			s.logger.Fatal().Err(err).Msg("")
		}
	}()

	return s.listenHttpService(context.Background())
}

func (s *incompletedServer) CompleteServer() error {
	return nil
}

func (s *incompletedServer) CustomGatewayMux(fn CustomGatewayMuxFunc) {
	s.fnCustomGatewayMux = fn
}

func (s *incompletedServer) listenGrpcService(gs *grpc.Server, ctx context.Context, port string) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return err
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			gs.GracefulStop()
			<-ctx.Done()
		}
	}()

	return gs.Serve(listen)
}

func (s *incompletedServer) listenHttpService(ctx context.Context) error {
	gwServer := &http.Server{
		Addr:              fmt.Sprintf(":%v", s.opts.HTTPGatewayPort),
		ReadHeaderTimeout: 5 * time.Second,
	}

	if s.fnCustomGatewayMux != nil {
		handler, err := s.fnCustomGatewayMux(s.gatewayMux)
		if err != nil {
			return err
		}
		gwServer.Handler = handler
	} else {
		gwServer.Handler = s.gatewayMux
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			if err := gwServer.Shutdown(ctx); err != nil {
				log.Fatalf("error shutdown gateway server: %v", err)
			}
			<-ctx.Done()
		}
	}()
	s.logger.Info().Msgf("Serving gRPC-Gateway on http://0.0.0.0:%v\n", s.opts.HTTPGatewayPort)
	return gwServer.ListenAndServe()
}
