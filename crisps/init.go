package crisps

import (
	"go.opentelemetry.io/otel/trace"
	"log"
	"time"
)

func InitProfiler(name string, version string) (err error) {
	log.Println("Initialising Profiler")
	var r = 10
	for i := 0; i < r; i++ {
		err = Profiler(name, version)
		if err == nil {
			log.Println("Profiler loaded")
			return nil
		} else {
			log.Println("Profiler failed: "+err.Error()+" (attempt ", i, " of ", r, ")")
			time.Sleep(2000 * time.Millisecond)
		}
	}
	return err
}
func InitTracing(name string, project string) (tracer trace.Tracer, err error) {
	log.Println("Initialising Tracing")
	var r = 10
	for i := 0; i < r; i++ {
		tracer, err = Tracer(name, project)
		if err == nil {
			log.Println("Tracing loaded")
			return tracer, nil
		} else {
			log.Println("Tracing failed: "+err.Error()+" (attempt ", i, " of ", r, ")")
			time.Sleep(2000 * time.Millisecond)
		}
	}
	return nil, err
}
