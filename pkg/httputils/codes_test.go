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
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHttpError(t *testing.T) {
	assert.True(t, IsHttpError(http.StatusBadRequest))          // error!
	assert.True(t, IsHttpError(http.StatusInternalServerError)) // error!
	assert.False(t, IsHttpError(http.StatusOK))                 // success
	assert.False(t, IsHttpError(http.StatusFound))              // redirect
	assert.False(t, IsHttpError(http.StatusSwitchingProtocols)) // informational
	assert.False(t, IsHttpError(601))                           // invalid code
}

func TestIsHttpInformation(t *testing.T) {
	assert.True(t, IsHttpInformation(http.StatusSwitchingProtocols))   // informational
	assert.False(t, IsHttpInformation(http.StatusOK))                  // success
	assert.False(t, IsHttpInformation(http.StatusFound))               // redirect
	assert.False(t, IsHttpInformation(http.StatusBadRequest))          // client error
	assert.False(t, IsHttpInformation(http.StatusInternalServerError)) // server error
	assert.False(t, IsHttpInformation(601))                            // invalid code
}

func TestIsHttpSuccess(t *testing.T) {
	assert.True(t, IsHttpSuccess(http.StatusOK))                  // success!
	assert.False(t, IsHttpSuccess(http.StatusBadRequest))         // error
	assert.False(t, IsHttpSuccess(http.StatusFound))              // redirect
	assert.False(t, IsHttpSuccess(http.StatusSwitchingProtocols)) // informational
	assert.False(t, IsHttpSuccess(601))                           // invalid code
}

func TestIsHttpRedirect(t *testing.T) {
	assert.False(t, IsHttpRedirect(http.StatusSwitchingProtocols))  // informational
	assert.False(t, IsHttpRedirect(http.StatusOK))                  // success
	assert.True(t, IsHttpRedirect(http.StatusFound))                // redirect
	assert.False(t, IsHttpRedirect(http.StatusBadRequest))          // client error
	assert.False(t, IsHttpRedirect(http.StatusInternalServerError)) // server error
	assert.False(t, IsHttpRedirect(601))                            // invalid code
}

func TestIsHttpClientError(t *testing.T) {
	assert.False(t, IsHttpClientError(http.StatusSwitchingProtocols))  // informational
	assert.False(t, IsHttpClientError(http.StatusOK))                  // success!
	assert.False(t, IsHttpClientError(http.StatusFound))               // redirect
	assert.True(t, IsHttpClientError(http.StatusBadRequest))           // client error
	assert.False(t, IsHttpClientError(http.StatusInternalServerError)) // server error
	assert.False(t, IsHttpClientError(601))                            // invalid code
}

func TestIsHttpServerError(t *testing.T) {
	assert.False(t, IsHttpServerError(http.StatusSwitchingProtocols)) // informational
	assert.False(t, IsHttpServerError(http.StatusOK))                 // success
	assert.False(t, IsHttpServerError(http.StatusFound))              // redirect
	assert.False(t, IsHttpServerError(http.StatusBadRequest))         // client error
	assert.True(t, IsHttpServerError(http.StatusInternalServerError)) // server error
	assert.False(t, IsHttpServerError(601))                           // invalid code
}
