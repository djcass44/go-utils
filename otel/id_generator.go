package otel

import (
	"context"
	"crypto/rand"
	"go.opentelemetry.io/otel/trace"
	"sync"
)

// IDGenerator generates trace/span IDs on demand.
// Largely borrowed from open-telemetry/opentelemetry-go which does not export
// their generator.
//
// https://github.com/open-telemetry/opentelemetry-go/blob/main/sdk/trace/id_generator.go
type IDGenerator struct {
	sync.Mutex
}

func (g *IDGenerator) NewIDs(context.Context) (trace.TraceID, trace.SpanID) {
	g.Lock()
	defer g.Unlock()

	tid := trace.TraceID{}
	_, _ = rand.Read(tid[:])
	sid := trace.SpanID{}
	_, _ = rand.Read(sid[:])
	return tid, sid
}

func (g *IDGenerator) NewSpanID(context.Context, trace.TraceID) trace.SpanID {
	g.Lock()
	defer g.Unlock()

	sid := trace.SpanID{}
	_, _ = rand.Read(sid[:])
	return sid
}
