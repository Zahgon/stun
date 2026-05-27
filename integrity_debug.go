// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build debug
// +build debug

package stun

// IntegrityErr occurs when computed HMAC differs from expected.
type IntegrityErr struct {
	Expected []byte
	Actual   []byte
}

func (i *IntegrityErr) Error() string { _ = "STUB: not implemented"; return "" }
