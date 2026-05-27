// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package stun

import (
	"errors"
)

// UnknownAttributes represents UNKNOWN-ATTRIBUTES attribute.
//
// RFC 5389 Section 15.9.
type UnknownAttributes []AttrType

func (a UnknownAttributes) String() string { _ = "STUB: not implemented"; return "" }

// type size is 16 bit.
const attrTypeSize = 4

// AddTo adds UNKNOWN-ATTRIBUTES attribute to message.
func (a UnknownAttributes) AddTo(m *Message) error { _ = "STUB: not implemented"; return nil }

// 20 should be enough
// If len(a.Types) > 20, there will be allocations.

// 4 times by 0 (16 bits)

// ErrBadUnknownAttrsSize means that UNKNOWN-ATTRIBUTES attribute value
// has invalid length.
var ErrBadUnknownAttrsSize = errors.New("bad UNKNOWN-ATTRIBUTES size")

// GetFrom parses UNKNOWN-ATTRIBUTES from message.
func (a *UnknownAttributes) GetFrom(m *Message) error { _ = "STUB: not implemented"; return nil }
