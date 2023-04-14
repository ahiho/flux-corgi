package sandbox

import (
	"context"

	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sandboxproto "fluxcorgi/pkg/proto/sandbox"
)

func (s *serviceImpl) DeploySandbox(ctx context.Context, req *sandboxproto.DeploySandboxRequest) (res *sandboxproto.DeploySandboxResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeploySandbox not implemented")
}
