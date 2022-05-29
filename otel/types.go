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

package otel

import (
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type Options struct {
	// Enabled determines whether OpenTelemetry integration should be activated on startup.
	Enabled bool
	// ServiceName is the name of the service that traces should be registered against.
	ServiceName string
	// Environment is a name used to differentiate different deployments (e.g. production or quality)
	Environment string
	// KubeNamespace is the kubernetes namespace that this pod is currently running in
	KubeNamespace string
	// SampleRate defines the fraction (0-1) of traces that should be sent to the exporter
	SampleRate float64
	// Exporter allows you to define a custom trace exporter.
	// Defaults to Jaeger if not provided.
	Exporter sdktrace.SpanExporter
}
