package swagger

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func HandleSwagger(token string) runtime.HandlerFunc {
	fs := http.FileServer(http.Dir("./third_party/swaggerui"))
	swaggerDocs, _ := os.ReadFile("./swagger/docs.swagger.json")

	swaggerHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/swagger/docs.swagger.json" {
			w.Header().Add("Content-Type", "application/json")
			if _, err := w.Write(swaggerDocs); err != nil {
				log.Fatalf("error when response swagger docs: %v", err)
			}
			return
		}
		http.StripPrefix("/swagger/", fs).ServeHTTP(w, r)
	}

	if os.Getenv("ENV") == "DEVELOPMENT" {
		return func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			swaggerHandler(w, r)
		}
	}

	return wrapBasicAuth(swaggerHandler, token)
}

func wrapBasicAuth(handler http.HandlerFunc, token string) runtime.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			writeBasicAuthFailed(w)
			return
		}
		if !(strings.HasPrefix(authHeader, "basic ") || strings.HasPrefix(authHeader, "Basic ")) {
			writeBasicAuthFailed(w)
			return
		}
		authToken := authHeader[len("basic "):]
		if token != authToken {
			writeBasicAuthFailed(w)
			return
		}

		handler(w, r)
	}
}

func writeBasicAuthFailed(w http.ResponseWriter) {
	w.Header().Add("WWW-Authenticate", `Basic realm="photogpt.com"`)
	w.WriteHeader(401)
}
