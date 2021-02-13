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
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

// WithBody reads an HTTP JSON request body and marshals it into a given struct
// Param v must be a pointer
func WithBody(r *http.Request, v interface{}) (err error) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, v)
	return
}

// ReturnJSON converts a given interface into JSON and writes it into an http response
// You should not write any additional data to the http.ResponseWriter after this
func ReturnJSON(w http.ResponseWriter, code int, v interface{}) {
	// convert our interface into JSON
	data, err := json.Marshal(v)
	if err != nil {
		log.WithError(err).Error("failed to marshal given struct into json")
		http.Error(w, "failed to marshal struct", http.StatusInternalServerError)
		return
	}
	// set the content type
	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(code)
	_, _ = w.Write(data)
}
