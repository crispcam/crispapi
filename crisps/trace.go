package crisps

import (
	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"log"
)

func Tracer(name string, project string) (trace.Tracer, error) {
	exporter, err := texporter.New(texporter.WithProjectID(project))
	if err != nil {
		log.Println("Failed to set up tracing: " + err.Error())
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exporter))
	otel.SetTracerProvider(tp)
	tracer := otel.GetTracerProvider().Tracer(name)
	return tracer, nil
}
