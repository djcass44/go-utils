package metrics

import (
	"context"
	"github.com/go-logr/logr"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/sdk/metric/aggregator/histogram"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	"go.opentelemetry.io/otel/sdk/metric/export/aggregation"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	selector "go.opentelemetry.io/otel/sdk/metric/selector/simple"
)

// New creates and configures an OpenTelemetry prometheus
// exporter.
// Providing a prometheus.Config is optional
func New(ctx context.Context, c *prometheus.Config, isGlobal bool) (*prometheus.Exporter, error) {
	log := logr.FromContextOrDiscard(ctx).WithName("prometheus")
	// ensure we have a config
	if c == nil {
		log.V(2).Info("consumer provided no prometheus config, using default")
		c = &prometheus.Config{
			DefaultHistogramBoundaries: []float64{1, 2, 5, 10, 20, 50},
		}
	}
	ctrl := controller.New(
		processor.NewFactory(
			selector.NewWithHistogramDistribution(
				histogram.WithExplicitBoundaries(c.DefaultHistogramBoundaries),
			),
			aggregation.CumulativeTemporalitySelector(),
			processor.WithMemory(true),
		),
	)

	exporter, err := prometheus.New(*c, ctrl)
	if err != nil {
		log.Error(err, "failed to create prometheus exporter")
		return nil, err
	}
	if isGlobal {
		log.V(2).Info("configuring exporter as global provider")
		global.SetMeterProvider(exporter.MeterProvider())
	}

	return exporter, nil
}
