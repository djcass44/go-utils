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

package sliceutils

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var strs = []string{"peach", "apple", "pear", "plum"}

func TestIndex(t *testing.T) {
	assert.Equal(t, 2, Index(strs, "pear"))
}

func TestInclude(t *testing.T) {
	assert.False(t, Includes(strs, "grape"))
}

func TestAny(t *testing.T) {
	assert.True(t, Any(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))
}

func TestAll(t *testing.T) {
	assert.False(t, All(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))
}

func TestFilter(t *testing.T) {
	assert.EqualValues(t, 3, len(Filter(strs, func(s string) bool {
		return strings.Contains(s, "e")
	})))
}

func TestMap(t *testing.T) {
	expected := []string{"PEACH", "APPLE", "PEAR", "PLUM"}
	assert.EqualValues(t, expected, Map(strs, strings.ToUpper))
}
