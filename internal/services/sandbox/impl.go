package sandbox

import (
	sandboxproto "fluxcorgi/pkg/proto/sandbox"
)

type serviceImpl struct {
	sandboxproto.UnimplementedSandboxServiceServer
}

func NewSandboxService() sandboxproto.SandboxServiceServer {
	return &serviceImpl{}
}
