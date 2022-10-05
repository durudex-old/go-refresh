/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package refresh_test

import (
	"testing"

	"github.com/durudex/go-refresh"
)

// Testing parsing refresh token string.
func Test_Parse(t *testing.T) {
	// Generating a new refresh token.
	r, err := refresh.New()
	if err != nil {
		t.Fatalf("error generating new refresh token: %s", err)
	}

	// Parsing refresh token string.
	token, err := refresh.Parse(r.Token("session", "object"))
	if err != nil {
		t.Fatalf("error parsing refresh token: %s", err)
	}

	// Checking is refresh token similar.
	if (token.Session == "session") != (token.Object == "object") {
		t.Fatal("error refresh token not similar")
	}
}
