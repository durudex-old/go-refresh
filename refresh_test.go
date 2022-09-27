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

// Testing generating a new refresh token.
func Test_New(t *testing.T) {
	// Generating a new refresh token.
	r, err := refresh.New()
	if err != nil {
		t.Fatalf("error generating new refresh token: %s", err)
	}

	// Checking refresh token is nil.
	if r.IsNil() {
		t.Fatal("error refresh token is nil")
	}
}

// Testing getting refresh token in string.
func TestToken_String(t *testing.T) {
	// Generating a new refresh token.
	r, err := refresh.New()
	if err != nil {
		t.Fatalf("error generating new refresh token: %s", err)
	}

	// Checking refresh token string is empty.
	if r.String() == "" {
		t.Fatal("error refresh token string is empty")
	}
}

// Testing hashing refresh token by secret key.
func TestToken_Hash(t *testing.T) {
	// Generating a new refresh token.
	r, err := refresh.New()
	if err != nil {
		t.Fatalf("error generating new refresh token: %s", err)
	}

	defer func() {
		if a := recover(); a != nil {
			t.Fatalf("error hashing token: %s", a)
		}
	}()

	// Hashing refresh token by secret key.
	r.Hash([]byte("durudex"))
}

// Testing parsing refresh token string.
func Test_Parse(t *testing.T) {
	// Generating a new refresh token.
	r, err := refresh.New()
	if err != nil {
		t.Fatalf("error generating new refresh token: %s", err)
	}

	// Parsing refresh token string.
	nr, id, err := refresh.Parse(r.Token("123"))
	if err != nil {
		t.Fatalf("error parsing refresh token: %s", err)
	}

	// Checking is refresh token similar.
	if (r.String() == nr.String()) != (id == "123") {
		t.Fatal("error refresh token not similar")
	}
}

// Testing getting refresh token from bytes.
func Test_FromBytes(t *testing.T) {
	// Generating a new refresh token.
	r, err := refresh.New()
	if err != nil {
		t.Fatalf("error generating new refresh token: %s", err)
	}

	// Getting refresh token from bytes.
	fb, err := refresh.FromBytes(r.Bytes())
	if err != nil {
		t.Fatalf("error getting refresh token from bytes: %s", err)
	}

	// Checking is refresh token bytes similar.
	if !reflect.DeepEqual(r.Bytes(), fb.Bytes()) {
		t.Fatal("error refresh token bytes not similar")
	}
}
