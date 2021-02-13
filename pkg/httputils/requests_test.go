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
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testStruct struct {
	Name string
	ID   int
}

func TestWithBody(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		var testInstance testStruct
		err := WithBody(r, &testInstance)
		if err != nil {
			t.Error(err)
		}
		_, _ = w.Write([]byte(testInstance.Name))
	}

	req := httptest.NewRequest(http.MethodPost, "http://example.org/test", strings.NewReader(`{"Name":"test","ID":54}`))
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "test", string(body))
}

func TestReturnJSON(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		ReturnJSON(w, http.StatusOK, &testStruct{
			Name: "test",
			ID:   54,
		})
	}

	req := httptest.NewRequest(http.MethodGet, "http://example.org/test", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, ApplicationJSON, resp.Header.Get(ContentType))
	assert.Equal(t, `{"Name":"test","ID":54}`, string(body))
}
