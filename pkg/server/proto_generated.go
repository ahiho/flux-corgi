package server

import (
	"context"

	sandboxproto "fluxcorgi/pkg/proto/sandbox"
)

func (s *incompletedServer) RegisterGateway(ctx context.Context) {
	sandboxproto.RegisterSandboxServiceHandler(ctx, s.gatewayMux, s.gatewayClientConn)
}
