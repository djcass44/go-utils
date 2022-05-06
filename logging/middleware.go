/*
 *    Copyright 2022 Django Cass
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 *
 */

package logging

import (
	"github.com/go-logr/logr"
	"go.opentelemetry.io/otel"
	"net/http"
)

const (
	KeyTraceID = "trace_id"
	KeySpanID  = "span_id"
)

type Middleware struct {
	log logr.Logger
}

func NewMiddleware(log logr.Logger) *Middleware {
	return &Middleware{
		log: log,
	}
}

func (m *Middleware) ServeHTTP(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, span := otel.Tracer("").Start(r.Context(), "logging_middleware")
		defer span.End()
		// create a new logger with our tracing information
		log := m.log.WithValues(
			KeyTraceID, span.SpanContext().TraceID(),
			KeySpanID, span.SpanContext().SpanID(),
		)
		// continue as normal
		h.ServeHTTP(w, r.WithContext(logr.NewContext(ctx, log)))
	})
}
