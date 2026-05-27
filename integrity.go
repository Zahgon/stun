// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package stun

import (
	//nolint:gosec
	//nolint:gosec
	"errors"
)

// separator for credentials.
const credentialsSep = ":"

// NewLongTermIntegrity returns new MessageIntegrity with key for long-term
// credentials. Password, username, and realm must be SASL-prepared.
func NewLongTermIntegrity(username, realm, password string) MessageIntegrity {
	_ = "STUB: not implemented"
	return *new(MessageIntegrity)
}

//nolint:gosec
//nolint:errcheck

// NewShortTermIntegrity returns new MessageIntegrity with key for short-term
// credentials. Password must be SASL-prepared.
func NewShortTermIntegrity(password string) MessageIntegrity {
	_ = "STUB: not implemented"
	return *new(MessageIntegrity)
}

// MessageIntegrity represents MESSAGE-INTEGRITY attribute.
//
// AddTo and Check methods are using zero-allocation version of hmac, see
// newHMAC function and internal/hmac/pool.go.
//
// RFC 5389 Section 15.4.
type MessageIntegrity []byte

func newHMAC(key, message, buf []byte) []byte { _ = "STUB: not implemented"; return nil }

func (i MessageIntegrity) String() string { _ = "STUB: not implemented"; return "" }

const messageIntegritySize = 20

// ErrFingerprintBeforeIntegrity means that FINGERPRINT attribute is already in
// message, so MESSAGE-INTEGRITY attribute cannot be added.
var ErrFingerprintBeforeIntegrity = errors.New("FINGERPRINT before MESSAGE-INTEGRITY attribute")

// AddTo adds MESSAGE-INTEGRITY attribute to message.
//
// CPU costly, see BenchmarkMessageIntegrity_AddTo.
func (i MessageIntegrity) AddTo(msg *Message) error { _ = "STUB: not implemented"; return nil }

// Message should not contain FINGERPRINT attribute
// before MESSAGE-INTEGRITY.

// The text used as input to HMAC is the STUN message,
// including the header, up to and including the attribute preceding the
// MESSAGE-INTEGRITY attribute.

// Adjusting m.Length to contain MESSAGE-INTEGRITY TLV.

// writing length to m.Raw
// calculating HMAC for adjusted m.Raw
// changing m.Length back

// Copy hmac value to temporary variable to protect it from resetting
// while processing m.Add call.

// ErrIntegrityMismatch means that computed HMAC differs from expected.
var ErrIntegrityMismatch = errors.New("integrity check failed")

// Check checks MESSAGE-INTEGRITY attribute.
//
// CPU costly, see BenchmarkMessageIntegrity_Check.
func (i MessageIntegrity) Check(msg *Message) error { _ = "STUB: not implemented"; return nil }

// Adjusting length in header to match m.Raw that was
// used when computing HMAC.

// /nolint:gosec

// startOfHMAC should be first byte of integrity attribute.

// data before integrity attribute

// writing length back
