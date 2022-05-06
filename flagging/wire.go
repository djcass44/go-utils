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
	"github.com/Unleash/unleash-client-go/v3"
	"github.com/go-logr/logr"
)

type Options struct {
	Token string
	Name  string
	URL   string
	Env   string
}

func Build(ctx context.Context, opt Options) {
	log := logr.FromContextOrDiscard(ctx).WithValues(
		"Name", opt.Name,
		"Env", opt.Env,
		"Url", opt.URL,
	)
	log.V(2).Info("preparing feature flag setup", "Token", opt.Token)
	if opt.Token == "" || opt.Name == "" || opt.URL == "" {
		log.V(1).Info("skipping feature flag setup")
		return
	}
	log.Info("enabling feature flags")
	opts := []unleash.ConfigOption{
		unleash.WithAppName(opt.Name),
		unleash.WithUrl(opt.URL),
		unleash.WithEnvironment(opt.Env),
		unleash.WithInstanceId(opt.Token),
	}
	if err := Initialize(opts...); err != nil {
		log.Error(err, "failed to initialise feature flag provider")
		return
	}
}
