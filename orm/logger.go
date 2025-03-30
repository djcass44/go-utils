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

package orm

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	"gorm.io/gorm/logger"
	"time"
)

type GormLogger struct {
	slowThreshold time.Duration
	log           logr.Logger
}

func NewGormLogger(log logr.Logger, slowThreshold time.Duration) *GormLogger {
	return &GormLogger{
		log:           log,
		slowThreshold: slowThreshold,
	}
}

func (g *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := g.log.V(int(level) - 1)
	return NewGormLogger(newLogger, g.slowThreshold)
}

func (g *GormLogger) Info(_ context.Context, s string, i ...interface{}) {
	g.log.Info(fmt.Sprintf(s, i...))
}

func (g *GormLogger) Warn(_ context.Context, s string, i ...interface{}) {
	g.log.Info(fmt.Sprintf(s, i...))
}

func (g *GormLogger) Error(_ context.Context, s string, i ...interface{}) {
	g.log.Info(fmt.Sprintf(s, i...))
}

func (g *GormLogger) Trace(_ context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	slow := elapsed > g.slowThreshold && g.slowThreshold != 0
	sql, rows := fc()
	kv := []any{"sql", sql, "rows", rows, "elapsed", elapsed, "slow", slow}
	if err != nil {
		g.log.V(4).Error(err, "tracing statement", kv...)
		return
	}
	g.log.V(4).Info("tracing statement", kv...)
}
