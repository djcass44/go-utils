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
	"gitlab.com/av1o/cap10/pkg/client"
	"go.opentelemetry.io/otel"
	"net/http"
)

const (
	KeyTraceID = "trace_id"
	KeySpanID  = "span_id"
	KeyUserSub = "sub"
	KeyUserIss = "iss"
)

func Middleware(rootLogger logr.Logger) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, span := otel.Tracer("").Start(r.Context(), "logging_middleware")
			defer span.End()
			// create a new logger with our tracing information
			log := rootLogger.WithValues(
				KeyTraceID, span.SpanContext().TraceID(),
				KeySpanID, span.SpanContext().SpanID(),
			)
			user, ok := client.GetContextUser(ctx)
			if ok {
				log = log.WithValues(
					KeyUserSub, user.Sub,
					KeyUserIss, user.Iss,
				)
			}
			// continue as normal
			handler.ServeHTTP(w, r.WithContext(logr.NewContext(ctx, log)))
		})
	}
}
