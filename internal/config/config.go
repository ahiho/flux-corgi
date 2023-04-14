package config

import (
	"context"

	"github.com/ahiho/gocandy/xcontext"
)

type ctxKeyType struct{}

var ctxKey = ctxKeyType{}

type Config struct {
	HTTPPort         string `env:"PORT" envDefault:"8080"`
	GrpcInternalPort string `env:"GRPC_INTERNAL_PORT" envDefault:"8082"`
	GrpcGatewayPort  string `env:"GRPC_GATEWAY_PORT" envDefault:"8081"`

	APIKey string `env:"API_KEY"`
}

func AppContextInjector(cfg *Config) xcontext.ValueBagInjector {
	return func(vb xcontext.ValueBag) {
		vb.AddValue(ctxKey, cfg)
	}
}

func FromContext(ctx context.Context) *Config {
	return ctx.Value(ctxKey).(*Config)
}
