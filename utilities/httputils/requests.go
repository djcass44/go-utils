/*
 *    Copyright 2020 Django Cass
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

package httputils

import (
	"context"
	"encoding/json"
	"github.com/go-logr/logr"
	"go.opentelemetry.io/otel"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"net/http"
)

// WithBody reads an HTTP JSON request body and marshals it into a given struct
// Param v must be a pointer
func WithBody(r *http.Request, v interface{}) error {
	ctx, span := otel.Tracer("").Start(r.Context(), "httputils_withBody")
	log := logr.FromContextOrDiscard(ctx)
	defer func() {
		log.V(2).Error(r.Body.Close(), "closing request body")
		span.End()
	}()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err, "failed to read request body")
		return err
	}
	log.V(2).Info("successfully read request body", "Body", string(body))
	return json.Unmarshal(body, v)
}

// WithProtoBody reads an HTTP JSON request body and marshals it into a
// given proto.Message
func WithProtoBody(r *http.Request, v proto.Message) error {
	ctx, span := otel.Tracer("").Start(r.Context(), "httputils_withProtoBody")
	log := logr.FromContextOrDiscard(ctx)
	defer func() {
		log.V(2).Error(r.Body.Close(), "closing request body")
		span.End()
	}()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err, "failed to read request body")
		return err
	}
	log.V(2).Info("successfully read request body", "Body", string(body))
	return protojson.Unmarshal(body, v)
}

// ReturnJSON converts a given interface into JSON and writes it into a http.ResponseWriter.
//
// You should not write any additional data to the http.ResponseWriter after this
func ReturnJSON(ctx context.Context, w http.ResponseWriter, code int, v interface{}) {
	ctx, span := otel.Tracer("").Start(ctx, "httputils_returnJson")
	defer span.End()
	log := logr.FromContextOrDiscard(ctx)
	// convert our interface into JSON
	data, err := json.Marshal(v)
	if err != nil {
		log.Error(err, "failed to marshal given struct into json")
		http.Error(w, "failed to marshal struct", http.StatusInternalServerError)
		return
	}
	log.V(2).Info("successfully marshalled JSON response", "Code", code, "Json", string(data))
	// set the content type
	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(code)
	_, _ = w.Write(data)
}

// ReturnProtoJSON converts a given proto.Message into JSON and writes
// it into an HTTP response.
//
// You should not write any additional data to the http.ResponseWriter
// after this.
func ReturnProtoJSON(ctx context.Context, w http.ResponseWriter, code int, v proto.Message) {
	ctx, span := otel.Tracer("").Start(ctx, "httputils_returnProtoJson")
	defer span.End()
	log := logr.FromContextOrDiscard(ctx)
	cfg := protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	data, err := cfg.Marshal(v)
	if err != nil {
		log.Error(err, "failed to marshal given proto.Message into JSON")
		http.Error(w, "failed to marshal response", http.StatusInternalServerError)
		return
	}
	log.V(2).Info("successfully marshalled JSON response", "Code", code, "Json", string(data))
	// set the content type
	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(code)
	_, _ = w.Write(data)
}
