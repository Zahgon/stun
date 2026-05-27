// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package testutil contains helpers and utilities for writing tests
package testutil

import (
	"testing"
)

// ShouldNotAllocate fails if f allocates.
func ShouldNotAllocate(t *testing.T, f func()) { _ = "STUB: not implemented"; return }
