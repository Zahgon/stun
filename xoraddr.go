// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package stun

import (
	"errors"
	"net"
)

const (
	familyIPv4 uint16 = 0x01
	familyIPv6 uint16 = 0x02
)

// XORMappedAddress implements XOR-MAPPED-ADDRESS attribute.
//
// RFC 5389 Section 15.2.
type XORMappedAddress struct {
	IP   net.IP
	Port int
}

func (a XORMappedAddress) String() string { _ = "STUB: not implemented"; return "" }

// isIPv4 returns true if ip with len of net.IPv6Len seems to be ipv4.
func isIPv4(ip net.IP) bool {
	_ = "STUB: not implemented"
	// Optimized for performance. Copied from net.IP.To4.
	return false
}

// Is p all zeros?
func isZeros(p net.IP) bool { _ = "STUB: not implemented"; return false }

// ErrBadIPLength means that len(IP) is not net.{IPv6len,IPv4len}.
var ErrBadIPLength = errors.New("invalid length of IP value")

// AddToAs adds XOR-MAPPED-ADDRESS value to msg as attr attribute.
func (a XORMappedAddress) AddToAs(msg *Message, attr AttrType) error {
	_ = "STUB: not implemented"
	return nil
}

// like in ip.To4()

// first 8 bits are zeroes

//nolint:gosec // G115, false positive, port

// AddTo adds XOR-MAPPED-ADDRESS to m. Can return ErrBadIPLength
// if len(a.IP) is invalid.
func (a XORMappedAddress) AddTo(m *Message) error { _ = "STUB: not implemented"; return nil }

// GetFromAs decodes XOR-MAPPED-ADDRESS attribute value in message
// getting it as for attr type.
func (a *XORMappedAddress) GetFromAs(msg *Message, attr AttrType) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensuring len(a.IP) == ipLen and reusing a.IP.

// GetFrom decodes XOR-MAPPED-ADDRESS attribute in message and returns
// error if any. While decoding, a.IP is reused if possible and can be
// rendered to invalid state (e.g. if a.IP was set to IPv6 and then
// IPv4 value were decoded into it), be careful.
//
// Example:
//
//	expectedIP := net.ParseIP("213.141.156.236")
//	expectedIP.String() // 213.141.156.236, 16 bytes, first 12 of them are zeroes
//	expectedPort := 21254
//	addr := &XORMappedAddress{
//	  IP:   expectedIP,
//	  Port: expectedPort,
//	}
//	// addr were added to message that is decoded as newMessage
//	// ...
//
//	addr.GetFrom(newMessage)
//	addr.IP.String()    // 213.141.156.236, net.IPv4Len
//	expectedIP.String() // d58d:9cec::ffff:d58d:9cec, 16 bytes, first 4 are IPv4
//	// now we have len(expectedIP) = 16 and len(addr.IP) = 4.
func (a *XORMappedAddress) GetFrom(m *Message) error { _ = "STUB: not implemented"; return nil }
