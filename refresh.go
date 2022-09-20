/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package refresh

import (
	"crypto/rand"
	"errors"
	"io"
	"strings"
	"sync"

	"github.com/jxskiss/base62"
)

// Refresh token bytes length.
const bytesLength = 24

var (
	// Refresh token point count error.
	ErrPointCount = errors.New("error token point count must be less one")
	// Refresh token payload size error.
	ErrPayloadSize = errors.New("error token payload size")

	// Nil refresh token value.
	Nil Token

	// Refresh token rander interface.
	rander = rand.Reader
	// Refresh token rand buffer.
	randBuffer = [bytesLength]byte{}
	// Refresh token rand sync mutex.
	randMutex = sync.Mutex{}
)

// Refresh token type.
type Token [bytesLength]byte

// Creating a new refresh token by id.
func New() (Token, error) {
	var token Token

	randMutex.Lock()

	// Generate random refresh token payload.
	_, err := io.ReadAtLeast(rander, randBuffer[:], len(randBuffer))
	copy(token[:], randBuffer[:])

	randMutex.Unlock()

	if err != nil {
		return Nil, err
	}

	return token, nil
}

// Getting refresh token in bytes.
func (t Token) Bytes() []byte {
	return t[:]
}

// Getting refresh token in string.
func (t Token) String() string {
	return base62.EncodeToString(t[:])
}

func (t Token) Token(id string) string {
	return id + "." + t.String()
}

// Checking refresh token is nil.
func (t Token) IsNil() bool {
	return t == Nil
}

// Parsing refresh token string.
func Parse(t string) (Token, string, error) {
	s := strings.Split(t, ".")

	// Checking count refresh token elements.
	if len(s) != 2 {
		return Nil, "", ErrPointCount
	}

	var token Token

	// Decode base62 payload string.
	b, err := base62.DecodeString(s[1])
	if err != nil {
		return Nil, "", err
	}

	// Checking refresh token payload bytes length.
	if len(b) != bytesLength {
		return Nil, "", ErrPayloadSize
	}

	copy(token[:], b[:])

	return token, s[0], nil
}
