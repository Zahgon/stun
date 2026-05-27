// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// This package implements RFC5780's tests:
// - 4.3.  Determining NAT Mapping Behavior
// - 4.4.  Determining NAT Filtering Behavior
package main

import (
	"errors"
	"flag"
	"net"
	"os"

	"github.com/pion/logging"
	"github.com/pion/stun/v3"
)

type stunServerConn struct {
	conn        net.PacketConn
	LocalAddr   net.Addr
	RemoteAddr  *net.UDPAddr
	OtherAddr   *net.UDPAddr
	messageChan chan *stun.Message
}

func (c *stunServerConn) Close() error { _ = "STUB: not implemented"; return nil }

var (
	//nolint:gochecknoglobals
	addrStrPtr = flag.String("server", "stun.voipgate.com:3478", "STUN server address")
	//nolint:gochecknoglobals
	timeoutPtr = flag.Int("timeout", 3, "the number of seconds to wait for STUN server's response")
	//nolint:gochecknoglobals
	verbose = flag.Int("verbose", 1, "the verbosity level")
	//nolint:gochecknoglobals
	log logging.LeveledLogger
)

const (
	messageHeaderSize = 20
)

var (
	errResponseMessage = errors.New("error reading from response message channel")
	errTimedOut        = errors.New("timed out waiting for response")
	errNoOtherAddress  = errors.New("no OTHER-ADDRESS in message")
)

func main() {
	flag.Parse()

	var logLevel logging.LogLevel
	switch *verbose {
	case 0:
		logLevel = logging.LogLevelWarn
	case 1:
		logLevel = logging.LogLevelInfo // default
	case 2:
		logLevel = logging.LogLevelDebug
	case 3:
		logLevel = logging.LogLevelTrace
	}
	log = logging.NewDefaultLeveledLoggerForScope("", logLevel, os.Stdout)

	if err := mappingTests(*addrStrPtr); err != nil {
		log.Warn("NAT mapping behavior: inconclusive")
	}
	if err := filteringTests(*addrStrPtr); err != nil {
		log.Warn("NAT filtering behavior: inconclusive")
	}
}

// RFC5780: 4.3.  Determining NAT Mapping Behavior.
func mappingTests(addrStr string) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

// Test I: Regular binding request

// Parse response message for XOR-MAPPED-ADDRESS and make sure OTHER-ADDRESS valid

// Assert mapping behavior

// Test II: Send binding request to the other address but primary port

// Assert mapping behavior

// Test III: Send binding request to the other address and port

// Assert mapping behavior

// RFC5780: 4.4.  Determining NAT Filtering Behavior.
func filteringTests(addrStr string) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

// Test I: Regular binding request

// Test II: Request to change both IP and port

// just to print out the resp

// something else went wrong

// Test III: Request to change port only

// just to print out the resp

// Parse a STUN message.
func parse(msg *stun.Message) (ret struct {
	xorAddr    *stun.XORMappedAddress
	otherAddr  *stun.OtherAddress
	respOrigin *stun.ResponseOrigin
	mappedAddr *stun.MappedAddress
	software   *stun.Software
},
) {
	_ = "STUB: not implemented"
	return nil
}

//nolint:staticcheck

// Given an address string, returns a StunServerConn.
func connect(addrStr string) (*stunServerConn, error) { _ = "STUB: not implemented"; return nil, nil }

// Send request and wait for response or timeout.
func (c *stunServerConn) roundTrip(msg *stun.Message, addr net.Addr) (*stun.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Wait for response or timeout

// taken from https://github.com/pion/stun/blob/master/cmd/stun-traversal/main.go
func listen(conn *net.UDPConn) (messages chan *stun.Message) { _ = "STUB: not implemented"; return nil }
