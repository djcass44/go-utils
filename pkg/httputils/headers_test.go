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
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var headers = &http.Header{
	"Content-Type": []string{ApplicationJson, ApplicationActuator},
	"X-My-Header":  []string{"Test"},
}

func TestHeadersAsMap(t *testing.T) {
	headers := HeadersAsMap(headers)
	expected := map[string][]string{
		"Content-Type": {ApplicationJson, ApplicationActuator},
		"X-My-Header":  {"Test"},
	}
	assert.EqualValues(t, expected, headers)
}

func TestHeadersAsFlatMap(t *testing.T) {
	headers := HeadersAsFlatMap(headers)
	expected := map[string]string{
		"Content-Type": ApplicationJson,
		"X-My-Header":  "Test",
	}
	assert.EqualValues(t, expected, headers)
}
