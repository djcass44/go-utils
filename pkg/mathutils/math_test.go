/*
 *    Copyright 2020 Django Cass
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

func TestMinFloat64(t *testing.T) {
	assert.Equal(t, 13.37, MinFloat64(13.37, 73.31))
	assert.Equal(t, 0.004, MinFloat64(15.423423, 0.004))
	assert.Equal(t, -3.0, MinFloat64(-3, -2.999999))
}

func TestMaxFloat64(t *testing.T) {
	assert.Equal(t, 73.31, MaxFloat64(13.37, 73.31))
	assert.Equal(t, 15.423423, MaxFloat64(15.423423, 0.004))
	assert.Equal(t, -2.999999, MaxFloat64(-3, -2.999999))
}
