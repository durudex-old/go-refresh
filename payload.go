/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package refresh

import (
	"crypto/rand"
	"fmt"
	"io"
	"sync"

	"github.com/jxskiss/base62"
	"golang.org/x/crypto/sha3"
)

// Refresh token payload bytes length.
const bytesLength = 24

var (
	// Refresh token payload nil value.
	PayloadNil Payload

	// Refresh token rander interface.
	rander = rand.Reader
	// Refresh token rand buffer.
	randBuffer = [bytesLength]byte{}
	// Refresh token rand sync mutex.
	randMutex = sync.Mutex{}
)

// Refresh token payload type.
type Payload [bytesLength]byte

// Generating a new refresh token payload.
func New() (Payload, error) {
	var payload Payload

	randMutex.Lock()

	// Generate random refresh token payload.
	_, err := io.ReadAtLeast(rander, randBuffer[:], len(randBuffer))
	copy(payload[:], randBuffer[:])

	randMutex.Unlock()

	if err != nil {
		return PayloadNil, err
	}

	return payload, nil
}

// Getting refresh token payload in bytes.
func (p Payload) Bytes() []byte {
	return p[:]
}

// Getting refresh token payload in string.
func (p Payload) String() string {
	return base62.EncodeToString(p[:])
}

// Getting full refresh token in string.
func (p Payload) Token(session, object string) string {
	return session + "." + object + "." + p.String()
}

// Checking refresh token payload is nil.
func (p Payload) IsNil() bool {
	return p == PayloadNil
}

// Hashing refresh token payload by secret key.
func (p Payload) Hash(secret []byte) []byte {
	h := sha3.New256()

	// Writing payload to hash.
	if _, err := h.Write(p[:]); err != nil {
		panic(fmt.Errorf("error writing payload bytes: %s", err))
	}
	// Writing secret key to hash.
	if _, err := h.Write(secret); err != nil {
		panic(fmt.Errorf("error writing secret bytes: %s", err))
	}

	return h.Sum(nil)
}

// Getting refresh token payload from bytes.
func FromBytes(b []byte) (Payload, error) {
	var token Payload

	// Checking refresh token payload bytes length.
	if len(b) != bytesLength {
		return PayloadNil, ErrPayloadSize
	}

	copy(token[:], b)

	return token, nil
}

// Getting refresh token payload from bytes. Same behavior as FromBytes, but
// returns a Nil Token on error.
func FromBytesOrNil(b []byte) Payload {
	token, err := FromBytes(b)
	if err != nil {
		return PayloadNil
	}

	return token
}

// Sets the global source rander of random bytes for refresh token payload generation.
func SetRand(r io.Reader) {
	if r == nil {
		return
	}

	rander = r
}
