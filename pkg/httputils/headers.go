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

import "net/http"

// collection of http headers
const (
	ContentType               = "Content-Type"
	AccessControlAllowOrigin  = "Access-Control-Allow-Origin"
	AccessControlAllowHeaders = "Access-Control-Allow-Headers"
	AccessControlAllowMethods = "Access-Control-Allow-Methods"
)

// HeadersAsMap converts an http.Header into a map
func HeadersAsMap(h *http.Header) map[string][]string {
	headers := map[string][]string{}
	for k, v := range *h {
		headers[k] = v
	}
	return headers
}

// HeadersAsFlatMap converts an http.Header into a map, flattening duplicate values
// e.g. X-My-Header=a,X-My-Header=b -> {X-My-Header: a}
func HeadersAsFlatMap(h *http.Header) map[string]string {
	headers := map[string]string{}
	for k := range *h {
		headers[k] = h.Get(k)
	}
	return headers
}

// RemoveHeaders removes all matching headers from an http.Header
func RemoveHeaders(h *http.Header, keys []string) {
	for _, k := range keys {
		h.Del(k)
	}
}
