package crisps

import (
	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
	"log"
)

func Tracer(name string, project string) (trace.Tracer, error) {
	exporter, err := texporter.New(texporter.WithProjectID(project))
	if err != nil {
		log.Println("Failed to set up tracing: " + err.Error())
		return nil, err
	}
	// Configure trace source information
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(name),
			semconv.ServiceNamespaceKey.String("crispcam"),
		),
	)
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(sdktrace.WithSyncer(exporter), sdktrace.WithResource(r))
	sdktrace.WithSampler(sdktrace.AlwaysSample())
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	log.Println("Started Tracer with name", name)
	return otel.GetTracerProvider().Tracer(name), nil
}
