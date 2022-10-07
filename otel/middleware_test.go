package otel

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/go-logr/logr/testr"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleware(t *testing.T) {
	ctx := logr.NewContext(context.TODO(), testr.NewWithOptions(t, testr.Options{Verbosity: 10}))
	mw := Middleware()
	assert.NotNil(t, mw)

	t.Run("no propagation generates a new trace ID", func(t *testing.T) {
		// verify that we are seeing the new trace id
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "https://example.org", nil)

		mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, span := otel.Tracer("").Start(r.Context(), "test")
			defer span.End()

			log := logr.FromContextOrDiscard(ctx)
			log.Info("test", "TraceID", span.SpanContext().TraceID(), "SpanID", span.SpanContext().SpanID())
		})).ServeHTTP(w, req.WithContext(ctx))
	})
}
