package otel

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/go-logr/logr/testr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetExporter(t *testing.T) {
	ctx := logr.NewContext(context.TODO(), testr.NewWithOptions(t, testr.Options{Verbosity: 10}))
	e, err := getExporter(ctx, &Options{})
	assert.NoError(t, err)
	assert.NotNil(t, e)
}
