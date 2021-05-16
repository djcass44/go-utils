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
	"bytes"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"errors"
	log "github.com/sirupsen/logrus"
)

var (
	// ErrMissingOrBadBlock is thrown when the PEM cannot be correctly decoded
	ErrMissingOrBadBlock = errors.New("failed to parse PEM block containing the key")
	// ErrKeyUnknown is thrown when the key type is not ed25519
	ErrKeyUnknown = errors.New("unexpected or unknown type of key")
)

// MarshalED25519PublicKey converts a crypto.ed25519 PublicKey into a string
func MarshalED25519PublicKey(pk *ed25519.PublicKey) (string, error) {
	asn1Bytes, err := x509.MarshalPKIXPublicKey(*pk)
	if err != nil {
		return "", err
	}
	pemKey := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}
	buf := new(bytes.Buffer)
	err = pem.Encode(buf, pemKey)
	if err != nil {
		return "", err
	}
	return buf.String(), err
}

// ParseED25519PrivateKey attempts to parse a crypto.ed25519 PrivateKey from a given byte slice
func ParseED25519PrivateKey(raw []byte) (*ed25519.PrivateKey, error) {
	// decode the pem
	block, _ := pem.Decode(raw)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, ErrMissingOrBadBlock
	}
	log.Debugf("successfully decoded PEM key of type %s", block.Type)
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// convert the resulting key into ed25519 or fail
	switch priv := priv.(type) {
	case ed25519.PrivateKey:
		return &priv, nil
	default:
		return nil, ErrKeyUnknown
	}
}

// ParseED25519PublicKey attempts to parse a crypto.ed25519 PublicKey from a given byte slice
func ParseED25519PublicKey(raw []byte) (*ed25519.PublicKey, error) {
	// decode the pem
	block, _ := pem.Decode(raw)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, ErrMissingOrBadBlock
	}
	log.Debugf("successfully decoded PEM key of type %s", block.Type)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// convert the resulting key into ed25519 or fail
	switch pub := pub.(type) {
	case ed25519.PublicKey:
		return &pub, nil
	default:
		return nil, ErrKeyUnknown
	}
}
