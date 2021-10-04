package crisps

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"net/http"
)

type contextKey string

var (
	contextKeyHeaders = contextKey("headers")
)

func (c contextKey) String() string {
	return "crisps context key: " + string(c)
}

func SendError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500"))
	fmt.Println(err.Error())
	return
}

// TraceRequest Persist Istio Tracing Headers
func TraceRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var headers = http.Header{}
		ctx := r.Context()
		tracingHeaders := []string{
			"x-request-id",
			"x-b3-traceid",
			"x-b3-spanid",
			"x-b3-sampled",
			"x-b3-parentspanid",
			"x-b3-flags",
			"x-ot-span-context",
		}
		for _, key := range tracingHeaders {
			if val := r.Header.Get(key); val != "" {
				// Persist headers for both GRPC and HTTP
				ctx = metadata.AppendToOutgoingContext(ctx, key, val)
				headers.Add(key, val)
			}
		}
		ctx = context.WithValue(ctx, contextKeyHeaders, headers)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

func TraceHeaders(ctx context.Context) (http.Header, bool) {
	headers, ok := ctx.Value(contextKeyHeaders).(http.Header)
	return headers, ok
}
