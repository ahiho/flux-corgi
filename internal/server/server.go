package server

import (
	"context"
	"encoding/base64"
	"net/http"
	"reflect"

	"github.com/ahiho/gocandy/grpcmdinjector"
	"github.com/ahiho/gocandy/requestid"
	"github.com/caarlos0/env"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/xid"
	"google.golang.org/grpc"

	"fluxcorgi/pkg/logger"
	sandboxproto "fluxcorgi/pkg/proto/sandbox"
	"fluxcorgi/pkg/server"
	"fluxcorgi/pkg/traceid"

	"fluxcorgi/internal/config"
	"fluxcorgi/internal/services/sandbox"
	"fluxcorgi/internal/swagger"
)

type serverApp struct {
	server.MonoServer
}

func RunServer() error {
	logger.Info().Msg("Load config")
	cfg := &config.Config{}
	if err := env.ParseWithFuncs(cfg, map[reflect.Type]env.ParserFunc{
		reflect.TypeOf([]byte{0}): func(v string) (interface{}, error) {
			return base64.StdEncoding.DecodeString(v)
		},
	}); err != nil {
		return err
	}
	opts := server.ServerOptions{
		Logger:           logger.Logger,
		GrpcGatewayPort:  cfg.GrpcGatewayPort,
		GrpcInternalPort: cfg.GrpcInternalPort,
		HTTPGatewayPort:  cfg.HTTPPort,
		GrpcGatewayOption: []grpc.ServerOption{
			grpc.ChainUnaryInterceptor(
				traceid.UnaryServerInterceptor(),
				logger.UnaryServerInterceptor(logger.ExtractMetadataField("x-request-id", "traceID")),
				// grpclogging.UnaryServerInterceptor(grpczerolog.InterceptorLogger(logger.Logger)),
				// grpcvalidator.UnaryServerInterceptor(true),
			),
		},
		GrpcInternalOption: []grpc.ServerOption{
			grpc.ChainUnaryInterceptor(
				traceid.UnaryServerInterceptor(),
				logger.UnaryServerInterceptor(logger.ExtractMetadataField("x-request-id", "traceID")),
				// grpclogging.UnaryServerInterceptor(grpczerolog.InterceptorLogger(logger.Logger)),
				// grpcvalidator.UnaryServerInterceptor(true),
			),
			grpc.ChainStreamInterceptor(
				traceid.StreamServerInterceptor(),
				logger.StreamServerInterceptor(logger.ExtractMetadataField("x-request-id", "traceID")),
				// grpclogging.StreamServerInterceptor(grpczerolog.InterceptorLogger(logger.Logger)),
				// grpcvalidator.StreamServerInterceptor(true),
			),
		},
	}

	logger.Info().Msg("Make new server")
	monoServer, err := server.NewIncompletedServer(opts)
	if err != nil {
		return err
	}

	sv := &serverApp{
		MonoServer: monoServer,
	}

	err = sv.CompleteServer()
	if err != nil {
		return err
	}
	err = sv.Run()
	return err
}

func (sa *serverApp) CompleteServer() error {
	sandboxService := sandbox.NewSandboxService()
	sa.RegisterGrpcService(&sandboxproto.SandboxService_ServiceDesc, sandboxService)
	sa.RegisterGateway(context.Background())

	sa.CustomGatewayMux(func(sm *runtime.ServeMux) (http.Handler, error) {
		if err := sm.HandlePath("GET", "/swagger/**", swagger.HandleSwagger("")); err != nil {
			return nil, err
		}

		if err := sm.HandlePath("GET", "/healthz", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			w.WriteHeader(http.StatusOK)
			if _, err := w.Write([]byte(`OK`)); err != nil {
				return
			}
		}); err != nil {
			return nil, err
		}

		return grpcmdinjector.WrapGrpcMD(
			requestid.InjectRequestID(
				sm, func() string {
					return xid.New().String()
				}),
		), nil
	})

	return nil
}
