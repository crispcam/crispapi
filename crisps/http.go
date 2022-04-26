package crisps

import (
	"context"
	"errors"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"google.golang.org/grpc/metadata"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type contextKey string

var (
	contextKeyHeaders = contextKey("headers")
)

func (c contextKey) String() string {
	return "crisps context key: " + string(c)
}

func Health(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = io.WriteString(w, `{"alive": true}`)
}

/*
Deprecated: Use SendErrorMsg() instead
*/
func SendError(w http.ResponseWriter, err error) {
	SendErrorMsg(w, err, "Error", http.StatusInternalServerError)
}

func SendErrorMsg(w http.ResponseWriter, err error, msg string, statusCode int) {
	http.Error(w, msg, statusCode)
	log.Println(err.Error())
}

func ServerError(w http.ResponseWriter, err error) {
	SendErrorMsg(w, err, "Internal Server Error", http.StatusInternalServerError)
}

// TraceRequest Persist Istio Tracing Headers
func TraceRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var headers = http.Header{}
		ctx := r.Context()
		tracingHeaders := []string{
			"user-agent",
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

func Request(ctx context.Context, u string, method string, form url.Values) ([]byte, error) {
	var result []byte
	req, err := http.NewRequestWithContext(ctx, method, u, strings.NewReader(form.Encode()))
	if err != nil {
		return result, err
	}

	// Assume json if there isn't a body, otherwise it's an encoded form
	if form == nil {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	} else {
		req.PostForm = form
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	if resp.StatusCode != http.StatusOK {
		return result, errors.New(fmt.Sprintf("upstream status code %d (request URI: %v)", resp.StatusCode, u))
	}
	result, err = ioutil.ReadAll(resp.Body)
	return result, nil
}
