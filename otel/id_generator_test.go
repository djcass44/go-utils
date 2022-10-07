package otel_test

import (
	"context"
	"github.com/djcass44/go-utils/otel"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/trace"
	"testing"
)

func TestIDGenerator_NewSpanID(t *testing.T) {
	g := new(otel.IDGenerator)
	sid := g.NewSpanID(context.TODO(), trace.TraceID{})
	assert.True(t, sid.IsValid())
}

func TestIDGenerator_NewIDs(t *testing.T) {
	g := new(otel.IDGenerator)
	tid, sid := g.NewIDs(context.TODO())
	assert.True(t, tid.IsValid())
	assert.True(t, sid.IsValid())
}
