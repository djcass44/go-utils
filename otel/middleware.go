package otel

import (
	"github.com/go-logr/logr"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

func Middleware() func(handler http.Handler) http.Handler {
	g := &IDGenerator{}
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, span := otel.Tracer("").Start(r.Context(), "tracing_middleware")
			defer span.End()
			log := logr.FromContextOrDiscard(ctx)

			if !span.SpanContext().HasTraceID() {
				log.V(6).Info("failed to locate trace ID - generating new span context")
				tid, sid := g.NewIDs(ctx)
				ctx = trace.ContextWithRemoteSpanContext(ctx, trace.NewSpanContext(trace.SpanContextConfig{
					TraceID: tid,
					SpanID:  sid,
					Remote:  false,
				}))
				log.V(6).Info("generated new trace", "NewTraceID", tid.String(), "NewSpanID", sid.String())
			}

			// continue as normal
			handler.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
