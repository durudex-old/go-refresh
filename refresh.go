/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package refresh

import (
	"errors"
	"strings"

	"github.com/jxskiss/base62"
)

var (
	// Refresh token point count error.
	ErrPointCount = errors.New("error token point count must be less one")
	// Refresh token payload size error.
	ErrPayloadSize = errors.New("error token payload size")

	// Refresh token nil value.
	Nil Token
)

// Refresh token structure.
type Token struct {
	// Refresh token session id.
	Session string
	// Refresh token object id.
	Object string
	// Refresh token payload.
	Payload Payload
}

// Parsing refresh token string.
func Parse(t string) (Token, error) {
	s := strings.Split(t, ".")

	// Checking count refresh token elements.
	if len(s) != 3 {
		return Nil, ErrPointCount
	}

	var token Token

	// Decode base62 payload string.
	b, err := base62.DecodeString(s[2])
	if err != nil {
		return Nil, err
	}

	// Checking refresh token payload bytes length.
	if len(b) != bytesLength {
		return Nil, ErrPayloadSize
	}

	token.Session = s[0]
	token.Object = s[1]

	copy(token.Payload[:], b[:])

	return token, nil
}
