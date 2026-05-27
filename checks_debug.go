// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build debug
// +build debug

package stun

// CheckSize returns *AttrLengthError if got is not equal to expected.
func CheckSize(a AttrType, got, expected int) error { _ = "STUB: not implemented"; return nil }

func checkHMAC(got, expected []byte) error { _ = "STUB: not implemented"; return nil }

func checkFingerprint(got, expected uint32) error { _ = "STUB: not implemented"; return nil }

// IsAttrSizeInvalid returns true if error means that attribute size is invalid.
func IsAttrSizeInvalid(err error) bool { _ = "STUB: not implemented"; return false }

// CheckOverflow returns *AttrOverflowErr if got is bigger that max.
func CheckOverflow(t AttrType, got, max int) error { _ = "STUB: not implemented"; return nil }

// IsAttrSizeOverflow returns true if error means that attribute size is too big.
func IsAttrSizeOverflow(err error) bool { _ = "STUB: not implemented"; return false }
