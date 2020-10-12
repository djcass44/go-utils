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

package cryptoutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var pemPublicKey = `-----BEGIN PUBLIC KEY-----
MCowBQYDK2VwAyEAwXHAsO+gSuBl61cxjndzjLOjFJS8JGD6IAxzADku708=
-----END PUBLIC KEY-----
`

func TestParseED25519PublicKeyWithValidKey(t *testing.T) {
	key, err := ParseED25519PublicKey([]byte(pemPublicKey))

	assert.NoError(t, err)
	assert.NotEmpty(t, key)
}

func TestParseED25519PublicKeyWithInvalidKey(t *testing.T) {
	key, err := ParseED25519PublicKey([]byte("this is totally not a key"))

	assert.Error(t, err)
	assert.Empty(t, key)
}

func TestMarshalED25519PublicKey(t *testing.T) {
	key, err := ParseED25519PublicKey([]byte(pemPublicKey))
	assert.NoError(t, err)
	assert.NotEmpty(t, key)

	str, err := MarshalED25519PublicKey(key)
	assert.NoError(t, err)
	assert.EqualValues(t, pemPublicKey, str)
}
