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
	metricDuration, _ = meter.SyncInt64().Histogram(
		"http.server.duration",
		instrument.WithUnit(unit.Milliseconds),
		instrument.WithDescription("Measures the duration of the inbound HTTP request."),
	)
	metricRequestSize, _ = meter.SyncInt64().Histogram(
		"http.server.request.size",
		instrument.WithUnit(unit.Bytes),
		instrument.WithDescription("Measures the size of the HTTP request."),
	)
	metricResponseSize, _ = meter.SyncInt64().Histogram(
		"http.server.response.size",
		instrument.WithUnit(unit.Bytes),
		instrument.WithDescription("Measures the size of the HTTP response."),
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
			metricRequestSize.Record(r.Context(), r.ContentLength, attributes...)
			metricActiveRequests.Add(r.Context(), 1, attributes...)
			// snoop the request
			m := httpsnoop.CaptureMetrics(handler, w, r)
			attributes = append(attributes, semconv.HTTPAttributesFromHTTPStatusCode(m.Code)...)
			// add our captured metrics
			metricDuration.Record(r.Context(), m.Duration.Milliseconds(), attributes...)
			metricResponseSize.Record(r.Context(), m.Written, attributes...)
			metricActiveRequests.Add(r.Context(), -1, attributes...)
		})
	}
}
