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
	"context"
	"github.com/go-logr/logr"
	"go.opentelemetry.io/otel/attribute"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type Options struct {
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
	Exporter *sdktrace.SpanExporter
}

// Build sets up the global OpenTelemetry context.
func Build(ctx context.Context, opts Options) error {
	log := logr.FromContextOrDiscard(ctx)
	log.V(1).Info("enabling OpenTelemetry", "Service", opts.ServiceName, "Env", opts.Environment, "k8s.namespace", opts.KubeNamespace, "SampleRate", opts.SampleRate, "DefaultExporter", opts.Exporter == nil)
	var exporter sdktrace.SpanExporter
	if opts.Exporter == nil {
		log.V(2).Info("using jaeger environment configuration", "Key", "OTEL_EXPORTER_JAEGER_ENDPOINT", "Value", os.Getenv("OTEL_EXPORTER_JAEGER_ENDPOINT"))
		log.V(3).Info("jaeger environment variables", "Endpoint", os.Getenv("OTEL_EXPORTER_JAEGER_ENDPOINT"), "User", os.Getenv("OTEL_EXPORTER_JAEGER_USER"), "Password", os.Getenv("OTEL_EXPORTER_JAEGER_PASSWORD"))
		e, err := jaeger.New(jaeger.WithCollectorEndpoint())
		if err != nil {
			log.Error(err, "failed to setup OpenTelemetry Jaeger exporter")
			return err
		}
		exporter = e
	}
	host, err := os.Hostname()
	if err != nil {
		log.Error(err, "failed to retrieve system hostname")
		host = "unknown"
	}
	log.V(1).Info("resolved hostname", "Hostname", host)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(opts.SampleRate)),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(opts.ServiceName),
			attribute.String("environment", opts.Environment),
			attribute.String("os", runtime.GOOS),
			attribute.String("arch", runtime.GOARCH),
			attribute.String("hostname", host),
			attribute.String("pod", host),
			attribute.String("namespace", opts.KubeNamespace),
			attribute.String("deployment.environment", opts.Environment),
			attribute.String("process.runtime.name", getRuntimeName()),
			attribute.String("process.runtime.version", runtime.Version()),
			// kubernetes semantic tags
			// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/resource/semantic_conventions/k8s.md
			attribute.String("k8s.pod.name", host),
			attribute.String("k8s.namespace.name", opts.KubeNamespace),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	go waitForShutdown(log, tp)

	return nil
}

func getRuntimeName() string {
	//goland:noinspection GoBoolExpressions
	if runtime.Compiler == "gc" {
		return "go"
	}
	return runtime.Compiler
}

func waitForShutdown(log logr.Logger, tp *sdktrace.TracerProvider) {
	log.V(1).Info("waiting for shutdown before closing OpenTelemetry provider")
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, syscall.SIGTERM, syscall.SIGINT)
	<-sigC
	if err := tp.Shutdown(context.Background()); err != nil {
		log.Error(err, "failed to shutdown OpenTelemetry tracer provider")
	}
}
