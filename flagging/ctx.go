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

package flagging

import (
	"context"
	flagctx "github.com/Unleash/unleash-client-go/v3/context"
	"github.com/go-logr/logr"
	"gitlab.com/av1o/cap10/pkg/client"
	"go.opentelemetry.io/otel"
)

func Context(parent context.Context) (ctx flagctx.Context) {
	parent, span := otel.Tracer("").Start(parent, "flagging_context")
	defer span.End()
	log := logr.FromContextOrDiscard(parent)
	user, ok := client.GetContextUser(parent)
	log.V(1).Info("checking for user in context", "Exists", ok)
	if !ok {
		return ctx
	}
	log.V(1).Info("successfully found user in context", "Sub", user.Sub, "Iss", user.Iss)
	ctx.UserId = user.AsUsername()
	ctx.Properties = user.Claims
	if ctx.Properties == nil {
		ctx.Properties = map[string]string{}
	}
	ctx.Properties["user_sub"] = user.Sub
	ctx.Properties["user_iss"] = user.Iss
	log.V(2).Info("set context properties", "Properties", ctx.Properties)
	return
}
