// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package stun

import (
	"errors"
)

// FingerprintAttr represents FINGERPRINT attribute.
//
// RFC 5389 Section 15.5.
type FingerprintAttr struct{}

// ErrFingerprintMismatch means that computed fingerprint differs from expected.
var ErrFingerprintMismatch = errors.New("fingerprint check failed")

// Fingerprint is shorthand for FingerprintAttr.
//
// Example:
//
//	m := New()
//	Fingerprint.AddTo(m)
var Fingerprint FingerprintAttr //nolint:gochecknoglobals

const (
	fingerprintXORValue uint32 = 0x5354554e //nolint:staticcheck
	fingerprintSize            = 4          // 32 bit
)

// FingerprintValue returns CRC-32 of b XOR-ed by 0x5354554e.
//
// The value of the attribute is computed as the CRC-32 of the STUN message
// up to (but excluding) the FINGERPRINT attribute itself, XOR'ed with
// the 32-bit value 0x5354554e (the XOR helps in cases where an
// application packet is also using CRC-32 in it).
func FingerprintValue(b []byte) uint32 { _ = "STUB: not implemented"; return 0 }

// XOR

// AddTo adds fingerprint to message.
func (FingerprintAttr) AddTo(m *Message) error {
	_ = "STUB: not implemented"

	// length in header should include size of fingerprint attribute
	return nil
}

// increasing length
// writing Length to Raw

// Check reads fingerprint value from m and checks it, returning error if any.
// Can return *AttrLengthErr, ErrAttributeNotFound, and *CRCMismatch.
func (FingerprintAttr) Check(m *Message) error { _ = "STUB: not implemented"; return nil }
