package metrics

import (
	"context"
	"github.com/go-logr/logr"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"os"
)

// New creates and configures an OpenTelemetry prometheus
// exporter.
// Providing a prometheus.Config is optional
func New(ctx context.Context, isGlobal bool) (*prometheus.Exporter, error) {
	log := logr.FromContextOrDiscard(ctx).WithName("prometheus")

	exporter, err := prometheus.New()
	if err != nil {
		log.Error(err, "failed to create prometheus exporter")
		return nil, err
	}
	if isGlobal {
		log.V(2).Info("configuring exporter as global provider")
		otel.SetMeterProvider(metric.NewMeterProvider(metric.WithReader(exporter)))
	}

	return exporter, nil
}

// NewDefault creates and configures OpenTelemetry for exporting
// Prometheus metrics. It uses the default Prometheus
// configuration and sets it as the global provider.
func NewDefault(ctx context.Context) (*prometheus.Exporter, error) {
	return New(ctx, true)
}

// MustNewDefault calls NewDefault and panics if there
// was an error.
func MustNewDefault(ctx context.Context) *prometheus.Exporter {
	log := logr.FromContextOrDiscard(ctx)
	prom, err := NewDefault(ctx)
	if err != nil {
		log.Error(err, "failed to configure OpenTelemetry Prometheus exporter")
		os.Exit(1)
		return nil
	}
	return prom
}
