// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package stun

import (
	"net"
)

// MappedAddress represents MAPPED-ADDRESS attribute.
//
// This attribute is used only by servers for achieving backwards
// compatibility with RFC 3489 clients.
//
// RFC 5389 Section 15.1.
type MappedAddress struct {
	IP   net.IP
	Port int
}

// AlternateServer represents ALTERNATE-SERVER attribute.
//
// RFC 5389 Section 15.11.
type AlternateServer struct {
	IP   net.IP
	Port int
}

// ResponseOrigin represents RESPONSE-ORIGIN attribute.
//
// RFC 5780 Section 7.3.
type ResponseOrigin struct {
	IP   net.IP
	Port int
}

// OtherAddress represents OTHER-ADDRESS attribute.
//
// RFC 5780 Section 7.4.
type OtherAddress struct {
	IP   net.IP
	Port int
}

// AddTo adds ALTERNATE-SERVER attribute to message.
func (s *AlternateServer) AddTo(m *Message) error { _ = "STUB: not implemented"; return nil }

// GetFrom decodes ALTERNATE-SERVER from message.
func (s *AlternateServer) GetFrom(m *Message) error { _ = "STUB: not implemented"; return nil }

func (a MappedAddress) String() string { _ = "STUB: not implemented"; return "" }

// GetFromAs decodes MAPPED-ADDRESS value in message m as an attribute of type t.
func (a *MappedAddress) GetFromAs(m *Message, t AttrType) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensuring len(a.IP) == ipLen and reusing a.IP.

// AddToAs adds MAPPED-ADDRESS value to m as t attribute.
func (a *MappedAddress) AddToAs(msg *Message, attrType AttrType) error {
	_ = "STUB: not implemented"
	return nil
}

// like in ip.To4()

//nolint:gosec //G115

// AddTo adds MAPPED-ADDRESS to message.
func (a *MappedAddress) AddTo(m *Message) error { _ = "STUB: not implemented"; return nil }

// GetFrom decodes MAPPED-ADDRESS from message.
func (a *MappedAddress) GetFrom(m *Message) error { _ = "STUB: not implemented"; return nil }

// AddTo adds OTHER-ADDRESS attribute to message.
func (o *OtherAddress) AddTo(m *Message) error { _ = "STUB: not implemented"; return nil }

// GetFrom decodes OTHER-ADDRESS from message.
func (o *OtherAddress) GetFrom(m *Message) error { _ = "STUB: not implemented"; return nil }

func (o OtherAddress) String() string { _ = "STUB: not implemented"; return "" }

// AddTo adds RESPONSE-ORIGIN attribute to message.
func (o *ResponseOrigin) AddTo(m *Message) error { _ = "STUB: not implemented"; return nil }

// GetFrom decodes RESPONSE-ORIGIN from message.
func (o *ResponseOrigin) GetFrom(m *Message) error { _ = "STUB: not implemented"; return nil }

func (o ResponseOrigin) String() string { _ = "STUB: not implemented"; return "" }
