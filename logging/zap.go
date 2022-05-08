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
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	stdlog "log"
)

// NewZap creates a logr.Logger using Zap.
// It fatally exits if the given configuration is invalid.
func NewZap(ctx context.Context, c zap.Config) (logr.Logger, context.Context) {
	z, err := c.Build()
	if err != nil {
		stdlog.Fatalf("failed to build zap logger configuration: %s", err)
		return logr.Discard(), nil
	}
	log := zapr.NewLogger(z)
	return log, logr.NewContext(ctx, log)
}
