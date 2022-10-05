/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package refresh_test

import (
	"reflect"
	"testing"

	"github.com/durudex/go-refresh"
)

// Testing generating a new refresh token payload.
func Test_New(t *testing.T) {
	// Generating a new refresh token payload.
	r, err := refresh.New()
	if err != nil {
		t.Fatalf("error generating new refresh token payload: %s", err)
	}

	// Checking refresh token is nil.
	if r.IsNil() {
		t.Fatal("error refresh token payload is nil")
	}
}

// Testing getting refresh token payload in string.
func TestPayload_String(t *testing.T) {
	// Generating a new refresh token payload.
	r, err := refresh.New()
	if err != nil {
		t.Fatalf("error generating new refresh token payload: %s", err)
	}

	// Checking refresh token payload string is empty.
	if r.String() == "" {
		t.Fatal("error refresh token payload string is empty")
	}
}

// Testing hashing refresh token payload by secret key.
func TestPayload_Hash(t *testing.T) {
	// Generating a new refresh token payload.
	r, err := refresh.New()
	if err != nil {
		t.Fatalf("error generating new refresh token payload: %s", err)
	}

	defer func() {
		if a := recover(); a != nil {
			t.Fatalf("error hashing refresh token payload: %s", a)
		}
	}()

	// Hashing refresh token payload by secret key.
	r.Hash([]byte("durudex"))
}

// Testing getting refresh token payload from bytes.
func Test_FromBytes(t *testing.T) {
	// Generating a new refresh token.
	r, err := refresh.New()
	if err != nil {
		t.Fatalf("error generating new refresh token: %s", err)
	}

	// Getting refresh token payload from bytes.
	p, err := refresh.FromBytes(r.Bytes())
	if err != nil {
		t.Fatalf("error getting refresh token from bytes: %s", err)
	}

	// Checking is refresh token bytes similar.
	if !reflect.DeepEqual(r.Bytes(), p.Bytes()) {
		t.Fatal("error refresh token bytes not similar")
	}
}
