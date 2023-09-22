package metrics

import (
	"github.com/felixge/httpsnoop"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/semconv/v1.20.0/httpconv"
	"net/http"
)

var (
	meter             = otel.GetMeterProvider().Meter("http")
	metricDuration, _ = meter.Int64Histogram(
		"http.server.duration",
		metric.WithUnit("{s}"),
		metric.WithDescription("Measures the duration of the inbound HTTP request."),
	)
	metricRequestSize, _ = meter.Int64Histogram(
		"http.server.request.size",
		metric.WithUnit("{By}"),
		metric.WithDescription("Measures the size of the HTTP request."),
	)
	metricResponseSize, _ = meter.Int64Histogram(
		"http.server.response.size",
		metric.WithUnit("{By}"),
		metric.WithDescription("Measures the size of the HTTP response."),
	)
	metricActiveRequests, _ = meter.Int64UpDownCounter(
		"http.server.active_requests",
		metric.WithUnit("{requests}"),
		metric.WithDescription("Measures the number of concurrent HTTP requests that are currently in-flight."),
	)
)

// Middleware collects prometheus metrics from HTTP
// server requests.
//
// See https://opentelemetry.io/docs/reference/specification/metrics/semantic_conventions/http-metrics/
// for exported metrics.
func Middleware() func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			serverName := r.Host
			// try to use the actual SNI
			// if available.
			if r.TLS != nil {
				serverName = r.TLS.ServerName
			}
			// collect standard attributes from the
			// incoming request
			attributes := httpconv.ServerRequest(serverName, r)
			metricRequestSize.Record(r.Context(), r.ContentLength, metric.WithAttributes(attributes...))
			metricActiveRequests.Add(r.Context(), 1, metric.WithAttributes(attributes...))
			// snoop the request
			m := httpsnoop.CaptureMetrics(handler, w, r)
			// add our captured metrics
			metricDuration.Record(r.Context(), m.Duration.Milliseconds(), metric.WithAttributes(attributes...))
			metricResponseSize.Record(r.Context(), m.Written, metric.WithAttributes(attributes...))
			metricActiveRequests.Add(r.Context(), -1, metric.WithAttributes(attributes...))
		})
	}
}
