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
	assert.True(t, IsHTTPError(http.StatusBadRequest))          // error!
	assert.True(t, IsHTTPError(http.StatusInternalServerError)) // error!
	assert.False(t, IsHTTPError(http.StatusOK))                 // success
	assert.False(t, IsHTTPError(http.StatusFound))              // redirect
	assert.False(t, IsHTTPError(http.StatusSwitchingProtocols)) // informational
	assert.False(t, IsHTTPError(601))                           // invalid code
}

func TestIsHttpInformation(t *testing.T) {
	assert.True(t, IsHTTPInformation(http.StatusSwitchingProtocols))   // informational
	assert.False(t, IsHTTPInformation(http.StatusOK))                  // success
	assert.False(t, IsHTTPInformation(http.StatusFound))               // redirect
	assert.False(t, IsHTTPInformation(http.StatusBadRequest))          // client error
	assert.False(t, IsHTTPInformation(http.StatusInternalServerError)) // server error
	assert.False(t, IsHTTPInformation(601))                            // invalid code
}

func TestIsHttpSuccess(t *testing.T) {
	assert.True(t, IsHTTPSuccess(http.StatusOK))                  // success!
	assert.False(t, IsHTTPSuccess(http.StatusBadRequest))         // error
	assert.False(t, IsHTTPSuccess(http.StatusFound))              // redirect
	assert.False(t, IsHTTPSuccess(http.StatusSwitchingProtocols)) // informational
	assert.False(t, IsHTTPSuccess(601))                           // invalid code
}

func TestIsHttpRedirect(t *testing.T) {
	assert.False(t, IsHTTPRedirect(http.StatusSwitchingProtocols))  // informational
	assert.False(t, IsHTTPRedirect(http.StatusOK))                  // success
	assert.True(t, IsHTTPRedirect(http.StatusFound))                // redirect
	assert.False(t, IsHTTPRedirect(http.StatusBadRequest))          // client error
	assert.False(t, IsHTTPRedirect(http.StatusInternalServerError)) // server error
	assert.False(t, IsHTTPRedirect(601))                            // invalid code
}

func TestIsHttpClientError(t *testing.T) {
	assert.False(t, IsHTTPClientError(http.StatusSwitchingProtocols))  // informational
	assert.False(t, IsHTTPClientError(http.StatusOK))                  // success!
	assert.False(t, IsHTTPClientError(http.StatusFound))               // redirect
	assert.True(t, IsHTTPClientError(http.StatusBadRequest))           // client error
	assert.False(t, IsHTTPClientError(http.StatusInternalServerError)) // server error
	assert.False(t, IsHTTPClientError(601))                            // invalid code
}

func TestIsHttpServerError(t *testing.T) {
	assert.False(t, IsHTTPServerError(http.StatusSwitchingProtocols)) // informational
	assert.False(t, IsHTTPServerError(http.StatusOK))                 // success
	assert.False(t, IsHTTPServerError(http.StatusFound))              // redirect
	assert.False(t, IsHTTPServerError(http.StatusBadRequest))         // client error
	assert.True(t, IsHTTPServerError(http.StatusInternalServerError)) // server error
	assert.False(t, IsHTTPServerError(601))                           // invalid code
}
