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

package flagging_test

import (
	"context"
	"github.com/djcass44/go-utils/flagging"
	"github.com/stretchr/testify/assert"
	"gitlab.com/av1o/cap10/pkg/client"
	"testing"
)

func TestContext(t *testing.T) {
	t.Run("valid user", func(t *testing.T) {
		ctx := client.PersistUserCtx(context.TODO(), nil, &client.UserClaim{
			Sub: "CN=Test User",
			Iss: "CN=Test Issuer",
			Claims: map[string]string{
				"Groups": "Foo, Bar",
			},
		})
		flagctx := flagging.Context(ctx)
		assert.NotNil(t, flagctx)
		assert.EqualValues(t, "CN=Test Issuer/CN=Test User", flagctx.UserId)
		assert.NotEmpty(t, flagctx.Properties)
		assert.EqualValues(t, "Foo, Bar", flagctx.Properties["Groups"])
		assert.Greater(t, len(flagctx.Properties), 2)
	})
	t.Run("no user", func(t *testing.T) {
		flagctx := flagging.Context(context.TODO())
		assert.NotNil(t, flagctx)
		assert.Nil(t, flagctx.Properties)
		assert.Empty(t, flagctx.UserId)
	})
	t.Run("user with no claims", func(t *testing.T) {
		ctx := client.PersistUserCtx(context.TODO(), nil, &client.UserClaim{
			Sub: "CN=Test User",
			Iss: "CN=Test Issuer",
		})
		flagctx := flagging.Context(ctx)
		assert.NotNil(t, flagctx)
		assert.EqualValues(t, "CN=Test Issuer/CN=Test User", flagctx.UserId)
		assert.NotNil(t, flagctx.Properties)
		assert.Len(t, flagctx.Properties, 2)
	})
}
