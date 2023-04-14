package sandbox

import (
	"context"

	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sandboxproto "fluxcorgi/pkg/proto/sandbox"
)

func (s *serviceImpl) ConfigSandbox(ctx context.Context, req *sandboxproto.ConfigSandboxRequest) (res *sandboxproto.ConfigSandboxResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigSandbox not implemented")
}
