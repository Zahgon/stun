// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package stuntest contains helpers for testing STUN clients
package stuntest

import (
	"errors"
	"net"
	"testing"
)

var errUDPServerUnsupportedNetwork = errors.New("unsupported network")

// NewUDPServer creates an udp server for testing.
// The supplied handler function will be called with the request
// and should be used to emulate the server behavior.
//
//nolint:cyclop
func NewUDPServer(
	t *testing.T,
	network string,
	maxMessageSize int,
	handler func(req []byte) ([]byte, error),
) (net.Addr, func(t *testing.T), error) {
	_ = "STUB: not implemented"
	return *new(net.Addr), nil, nil
}

// Necessary for IPv6
//nolint:forcetypeassert
