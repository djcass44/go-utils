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
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestNewZap(t *testing.T) {
	var cases = []struct {
		name string
		in   zap.Config
	}{
		{
			"production config",
			zap.NewProductionConfig(),
		},
		{
			"development config",
			zap.NewDevelopmentConfig(),
		},
		{
			"custom config",
			zap.Config{Encoding: "json", Level: zap.NewAtomicLevelAt(zapcore.Level(-2))},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			log, ctx := NewZap(context.TODO(), tt.in)
			assert.NotNil(t, log)
			assert.NotNil(t, ctx)
		})
	}
}
