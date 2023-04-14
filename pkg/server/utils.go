package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	ErrBadRequest   = "ERR_BAD_REQUEST"
	ErrUnauthorized = "ERR_UNAUTHORIZED"
	ErrForbidden    = "ERR_FORBIDDEN"
	ErrNotFound     = "ERR_NOTFOUND"
	ErrInternal     = "ERR_INTERNAL"
)

var (
	AllowedHTTPErrorStatuses = []int{400, 401, 403, 404}
)

func isHTTPCodeAllowed(statusCode int) bool {
	for _, val := range AllowedHTTPErrorStatuses {
		if statusCode == val {
			return true
		}
	}
	return false
}

type errorStruct struct {
	Code    string `json:"code"`
	Message string `json:"message,omitempty"`
	TraceID string `json:"traceId,omitempty"`
}

func defaultCode(code codes.Code) string {
	switch code {
	case codes.InvalidArgument:
		return ErrBadRequest
	case codes.Unauthenticated:
		return ErrUnauthorized
	case codes.PermissionDenied:
		return ErrForbidden
	case codes.NotFound:
		return ErrNotFound
	default:
		return ErrInternal
	}
}

func FormatRestError(showInternalError bool) runtime.ErrorHandlerFunc {
	return func(ctx context.Context, sm *runtime.ServeMux, m runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
		w.Header().Add("Content-Type", "application/json")
		grpcErr := status.Convert(err)

		statusCode := runtime.HTTPStatusFromCode(grpcErr.Code())
		if !isHTTPCodeAllowed(statusCode) {
			statusCode = http.StatusInternalServerError
		}

		e := &errorStruct{
			Code: defaultCode(grpcErr.Code()),
		}

		e.TraceID = r.Header.Get("Grpc-Metadata-X-Request-Id")
		if statusCode != 500 || showInternalError {
			e.Message = grpcErr.Message()
		}

		w.WriteHeader(statusCode)
		bytes, _ := json.Marshal(e)
		_, _ = w.Write(bytes)
	}
}
