package mathutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	assert.Equal(t, 1, Min(1, 10))
	assert.Equal(t, 0, Min(0, 1))
	assert.Equal(t, -1, Min(0, -1))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 10, Max(1, 10))
	assert.Equal(t, 1, Max(0, 1))
	assert.Equal(t, 0, Max(0, -1))
}
