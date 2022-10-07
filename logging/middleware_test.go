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
	"context"
	"github.com/go-logr/logr"
	"github.com/go-logr/logr/testr"
	"github.com/stretchr/testify/assert"
	"gitlab.com/av1o/cap10/pkg/client"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewMiddleware(t *testing.T) {
	ctx := logr.NewContext(context.TODO(), testr.NewWithOptions(t, testr.Options{Verbosity: 10}))
	mw := Middleware(logr.FromContextOrDiscard(ctx))
	assert.NotNil(t, mw)

	// verify that we are seeing the log statements
	// that we expect
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "https://example.org", nil)

	userCtx := client.PersistUserCtx(ctx, nil, &client.UserClaim{
		Sub: "CN=Test User",
		Iss: "CN=Test Issuer",
	})

	mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logr.FromContextOrDiscard(r.Context())
		log.Info("test")
	})).ServeHTTP(w, req.WithContext(userCtx))
}
