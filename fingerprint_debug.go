// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build debug
// +build debug

package stun

// CRCMismatch represents CRC check error.
type CRCMismatch struct {
	Expected uint32
	Actual   uint32
}

func (m CRCMismatch) Error() string { _ = "STUB: not implemented"; return "" }
