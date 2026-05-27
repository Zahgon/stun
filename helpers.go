// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package stun

// Interfaces that are implemented by message attributes, shorthands for them,
// or helpers for message fields as type or transaction id.
type (
	// Setter sets *Message attribute.
	Setter interface {
		AddTo(m *Message) error
	}
	// Getter parses attribute from *Message.
	Getter interface {
		GetFrom(m *Message) error
	}
	// Checker checks *Message attribute.
	Checker interface {
		Check(m *Message) error
	}
)

// Build resets message and applies setters to it in batch, returning on
// first error. To prevent allocations, pass pointers to values.
//
// Example:
//
//	var (
//		t        = BindingRequest
//		username = NewUsername("username")
//		nonce    = NewNonce("nonce")
//		realm    = NewRealm("example.org")
//	)
//	m := new(Message)
//	m.Build(t, username, nonce, realm)     // 4 allocations
//	m.Build(&t, &username, &nonce, &realm) // 0 allocations
//
// See BenchmarkBuildOverhead.
func (m *Message) Build(setters ...Setter) error { _ = "STUB: not implemented"; return nil }

// Check applies checkers to message in batch, returning on first error.
func (m *Message) Check(checkers ...Checker) error { _ = "STUB: not implemented"; return nil }

// Parse applies getters to message in batch, returning on first error.
func (m *Message) Parse(getters ...Getter) error { _ = "STUB: not implemented"; return nil }

// MustBuild wraps Build call and panics on error.
func MustBuild(setters ...Setter) *Message { _ = "STUB: not implemented"; return nil }

//nolint

// Build wraps Message.Build method.
func Build(setters ...Setter) (*Message, error) { _ = "STUB: not implemented"; return nil, nil }

// ForEach is helper that iterates over message attributes allowing to call
// Getter in f callback to get all attributes of type t and returning on first
// f error.
//
// The m.Get method inside f will be returning next attribute on each f call.
// Does not error if there are no results.
func (m *Message) ForEach(t AttrType, f func(m *Message) error) error {
	_ = "STUB: not implemented"
	return nil
}
