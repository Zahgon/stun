// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build !debug

package stun

// CheckSize returns ErrAttrSizeInvalid if got is not equal to expected.
func CheckSize(_ AttrType, got, expected int) error { _ = "STUB: not implemented"; return nil }

func checkHMAC(got, expected []byte) error { _ = "STUB: not implemented"; return nil }

func checkFingerprint(got, expected uint32) error { _ = "STUB: not implemented"; return nil }

// IsAttrSizeInvalid returns true if error means that attribute size is invalid.
func IsAttrSizeInvalid(err error) bool { _ = "STUB: not implemented"; return false }

// CheckOverflow returns ErrAttributeSizeOverflow if got is bigger that max.
func CheckOverflow(_ AttrType, got, maxVal int) error { _ = "STUB: not implemented"; return nil }

// IsAttrSizeOverflow returns true if error means that attribute size is too big.
func IsAttrSizeOverflow(err error) bool { _ = "STUB: not implemented"; return false }
