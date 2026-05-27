// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package stun

// NewUsername returns Username with provided value.
func NewUsername(username string) Username {
	_ = "STUB: not implemented"
	return *

	// Username represents USERNAME attribute.
	//
	// RFC 5389 Section 15.3.
	new(Username)
}

type Username []byte

func (u Username) String() string { _ = "STUB: not implemented"; return "" }

const maxUsernameB = 513

// AddTo adds USERNAME attribute to message.
func (u Username) AddTo(m *Message) error { _ = "STUB: not implemented"; return nil }

// GetFrom gets USERNAME from message.
func (u *Username) GetFrom(m *Message) error { _ = "STUB: not implemented"; return nil }

// NewRealm returns Realm with provided value.
// Must be SASL-prepared.
func NewRealm(realm string) Realm {
	_ = "STUB: not implemented"
	return *

	// Realm represents REALM attribute.
	//
	// RFC 5389 Section 15.7.
	new(Realm)
}

type Realm []byte

func (n Realm) String() string { _ = "STUB: not implemented"; return "" }

const maxRealmB = 763

// AddTo adds NONCE to message.
func (n Realm) AddTo(m *Message) error { _ = "STUB: not implemented"; return nil }

// GetFrom gets REALM from message.
func (n *Realm) GetFrom(m *Message) error { _ = "STUB: not implemented"; return nil }

const softwareRawMaxB = 763

// Software is SOFTWARE attribute.
//
// RFC 5389 Section 15.10.
type Software []byte

func (s Software) String() string {
	_ = "STUB: not implemented"

	// NewSoftware returns *Software from string.
	return ""
}

func NewSoftware(software string) Software {
	_ = "STUB: not implemented"
	return *

	// AddTo adds Software attribute to m.
	new(Software)
}

func (s Software) AddTo(m *Message) error { _ = "STUB: not implemented"; return nil }

// GetFrom decodes Software from m.
func (s *Software) GetFrom(m *Message) error { _ = "STUB: not implemented"; return nil }

// Nonce represents NONCE attribute.
//
// RFC 5389 Section 15.8.
type Nonce []byte

// NewNonce returns new Nonce from string.
func NewNonce(nonce string) Nonce { _ = "STUB: not implemented"; return *new(Nonce) }

func (n Nonce) String() string { _ = "STUB: not implemented"; return "" }

const maxNonceB = 763

// AddTo adds NONCE to message.
func (n Nonce) AddTo(m *Message) error { _ = "STUB: not implemented"; return nil }

// GetFrom gets NONCE from message.
func (n *Nonce) GetFrom(m *Message) error { _ = "STUB: not implemented"; return nil }

// TextAttribute is helper for adding and getting text attributes.
type TextAttribute []byte

// AddToAs adds attribute with type t to m, checking maximum length. If maxLen
// is less than 0, no check is performed.
func (v TextAttribute) AddToAs(m *Message, t AttrType, maxLen int) error {
	_ = "STUB: not implemented"
	return nil
}

// GetFromAs gets t attribute from m and appends its value to reseted v.
func (v *TextAttribute) GetFromAs(m *Message, t AttrType) error {
	_ = "STUB: not implemented"
	return nil
}
