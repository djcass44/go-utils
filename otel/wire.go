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
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

// Build sets up the global OpenTelemetry context.
func Build(ctx context.Context, opts Options) error {
	log := logr.FromContextOrDiscard(ctx)

	// skip if not enabled
	if !opts.Enabled {
		log.V(1).Info("skipping OpenTelemetry integration as it's currently disabled")
		return nil
	}

	// if no namespace was provided, check
	// the environment
	if opts.KubeNamespace == "" {
		opts.KubeNamespace = os.Getenv("KUBE_NAMESPACE")
	}

	log.V(1).Info("enabling OpenTelemetry", "Service", opts.ServiceName, "Env", opts.Environment, "k8s.namespace", opts.KubeNamespace, "SampleRate", opts.SampleRate, "DefaultExporter", opts.Exporter == nil)
	exporter, err := getExporter(ctx, &opts)
	if err != nil {
		return err
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

func getExporter(ctx context.Context, opts *Options) (sdktrace.SpanExporter, error) {
	log := logr.FromContextOrDiscard(ctx)
	// if no explicit exporter is provided, default
	// to the Jaeger exporter
	if opts.Exporter == nil {
		log.V(2).Info("using jaeger environment configuration", "Key", "OTEL_EXPORTER_JAEGER_ENDPOINT", "Value", os.Getenv("OTEL_EXPORTER_JAEGER_ENDPOINT"))
		log.V(3).Info("jaeger environment variables", "Endpoint", os.Getenv("OTEL_EXPORTER_JAEGER_ENDPOINT"), "User", os.Getenv("OTEL_EXPORTER_JAEGER_USER"), "Password", os.Getenv("OTEL_EXPORTER_JAEGER_PASSWORD"))
		e, err := jaeger.New(jaeger.WithCollectorEndpoint())
		if err != nil {
			log.Error(err, "failed to setup OpenTelemetry Jaeger exporter")
			return nil, err
		}
		return e, nil
	}
	log.V(2).Info("using consumer-provided exporter")
	return opts.Exporter, nil
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
