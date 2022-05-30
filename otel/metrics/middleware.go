package metrics

import (
	"github.com/felixge/httpsnoop"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/metric/unit"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"net/http"
)

var (
	meter             = global.MeterProvider().Meter("http")
	metricDuration, _ = meter.SyncInt64().Counter(
		"http.server.duration",
		instrument.WithUnit(unit.Milliseconds),
		instrument.WithDescription("Measures the duration of the inbound HTTP request."),
	)
	metricActiveRequests, _ = meter.SyncInt64().UpDownCounter(
		"http.server.active_requests",
		instrument.WithUnit("{requests}"),
		instrument.WithDescription("Measures the number of concurrent HTTP requests that are currently in-flight."),
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
			attributes := semconv.HTTPServerMetricAttributesFromHTTPRequest(serverName, r)
			metricActiveRequests.Add(r.Context(), 1, attributes...)
			// snoop the request
			m := httpsnoop.CaptureMetrics(handler, w, r)
			// add our captured metrics
			metricDuration.Add(r.Context(), m.Duration.Milliseconds(), attributes...)
			metricActiveRequests.Add(r.Context(), -1, attributes...)
		})
	}
}
