package crisps

import (
	"cloud.google.com/go/profiler"
)

func Profiler(service string, version string) error {
	cfg := profiler.Config{
		Service:           service,
		ServiceVersion:    version,
		EnableOCTelemetry: true,
	}

	// Profiler initialization, best done as early as possible.
	if err := profiler.Start(cfg); err != nil {
		return err
	}
	return nil
}
