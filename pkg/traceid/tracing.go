package traceid

import (
	"context"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	metautils "github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type traceIDCtxKeyType struct{}

var ctxKey = traceIDCtxKeyType{}

const defaultTraceID = "NO_TRACE_ID"

func AnnotateTraceID(ctx context.Context) context.Context {
	md := metautils.ExtractIncoming(ctx)

	// var traceID string
	traceID := md.Get("x-trace-id")
	if traceID == "" {
		traceID = md.Get("x-request-id")
	}

	if traceID == "" {
		traceID = defaultTraceID
	}

	ctx2 := context.WithValue(ctx, ctxKey, traceID)

	return metadata.AppendToOutgoingContext(ctx2, "x-trace-id", traceID, "x-request-id", traceID)
}

func Clone(ctx context.Context) context.Context {
	newCtx := context.Background()
	traceID := FromContext(ctx)
	newCtx = context.WithValue(newCtx, ctxKey, traceID)
	return metadata.AppendToOutgoingContext(newCtx, "x-trace-id", traceID, "x-request-id", traceID)
}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		return handler(AnnotateTraceID(ctx), req)
	}
}

func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		wrapped := middleware.WrapServerStream(ss)
		wrapped.WrappedContext = AnnotateTraceID(wrapped.WrappedContext)
		return handler(srv, wrapped)
	}
}

func FromContext(ctx context.Context) string {
	val := ctx.Value(ctxKey)
	if val == nil {
		return defaultTraceID
	}
	traceID, ok := val.(string)
	if !ok {
		return defaultTraceID
	}
	return traceID
}
