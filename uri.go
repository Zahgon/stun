// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package stun

import (
	"errors"
)

var (
	// ErrUnknownType indicates an error with Unknown info.
	ErrUnknownType = errors.New("Unknown")

	// ErrSchemeType indicates the scheme type could not be parsed.
	ErrSchemeType = errors.New("unknown scheme type")

	// ErrSTUNQuery indicates query arguments are provided in a STUN URL.
	ErrSTUNQuery = errors.New("queries not supported in stun address")

	// ErrInvalidQuery indicates an malformed query is provided.
	ErrInvalidQuery = errors.New("invalid query")

	// ErrHost indicates malformed hostname is provided.
	ErrHost = errors.New("invalid hostname")

	// ErrPort indicates malformed port is provided.
	ErrPort = errors.New("invalid port")

	// ErrProtoType indicates an unsupported transport type was provided.
	ErrProtoType = errors.New("invalid transport protocol type")
)

// SchemeType indicates the type of server used in the ice.URL structure.
type SchemeType int

const (
	// SchemeTypeUnknown indicates an unknown or unsupported scheme.
	SchemeTypeUnknown SchemeType = iota

	// SchemeTypeSTUN indicates the URL represents a STUN server.
	SchemeTypeSTUN

	// SchemeTypeSTUNS indicates the URL represents a STUNS (secure) server.
	SchemeTypeSTUNS

	// SchemeTypeTURN indicates the URL represents a TURN server.
	SchemeTypeTURN

	// SchemeTypeTURNS indicates the URL represents a TURNS (secure) server.
	SchemeTypeTURNS
)

// NewSchemeType defines a procedure for creating a new SchemeType from a raw
// string naming the scheme type.
func NewSchemeType(raw string) SchemeType { _ = "STUB: not implemented"; return *new(SchemeType) }

func (t SchemeType) String() string { _ = "STUB: not implemented"; return "" }

// ProtoType indicates the transport protocol type that is used in the ice.URL
// structure.
type ProtoType int

const (
	// ProtoTypeUnknown indicates an unknown or unsupported protocol.
	ProtoTypeUnknown ProtoType = iota

	// ProtoTypeUDP indicates the URL uses a UDP transport.
	ProtoTypeUDP

	// ProtoTypeTCP indicates the URL uses a TCP transport.
	ProtoTypeTCP
)

// NewProtoType defines a procedure for creating a new ProtoType from a raw
// string naming the transport protocol type.
func NewProtoType(raw string) ProtoType {
	_ = "STUB: not implemented"
	return *

	//nolint:goconst
	new(ProtoType)
}

//nolint:goconst

func (t ProtoType) String() string { _ = "STUB: not implemented"; return "" }

//nolint:goconst

// URI represents a STUN (rfc7064) or TURN (rfc7065) URI.
type URI struct {
	Scheme   SchemeType
	Host     string
	Port     int
	Username string
	Password string //nolint:gosec // G117 -- no hardcoded credentials.
	Proto    ProtoType
}

// ParseURI parses a STUN or TURN urls following the ABNF syntax described in
// https://tools.ietf.org/html/rfc7064 and https://tools.ietf.org/html/rfc7065
// respectively.
func ParseURI(raw string) (*URI, error) {
	_ = "STUB: not implemented" //nolint:gocognit,cyclop
	return nil, nil
}

//nolint:nestif

func parseProto(raw string) (ProtoType, error) {
	_ = "STUB: not implemented"
	return *new(ProtoType), nil
}

func (u URI) String() string { _ = "STUB: not implemented"; return "" }

// IsSecure returns whether the this URL's scheme describes secure scheme or not.
func (u URI) IsSecure() bool { _ = "STUB: not implemented"; return false }
