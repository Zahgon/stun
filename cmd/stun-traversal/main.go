// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package main implements a simple CLI tools to perform NAT traversal via STUN
package main

import (
	"flag"
	"log"
	"net"
	"time"

	"github.com/pion/stun/v3"
)

var server = flag.String("server", "stun.voipgate.com:3478", "Stun server address") //nolint:gochecknoglobals

const (
	udp           = "udp4"
	pingMsg       = "ping"
	pongMsg       = "pong"
	timeoutMillis = 500
)

func main() { //nolint:gocognit,cyclop
	flag.Parse()

	srvAddr, err := net.ResolveUDPAddr(udp, *server)
	if err != nil {
		log.Fatalf("Failed to resolve server addr: %s", err)
	}

	conn, err := net.ListenUDP(udp, nil)
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}

	defer func() {
		_ = conn.Close()
	}()

	log.Printf("Listening on %s", conn.LocalAddr())

	var publicAddr stun.XORMappedAddress
	var peerAddr *net.UDPAddr

	messageChan := listen(conn)
	var peerAddrChan <-chan string

	keepalive := time.Tick(timeoutMillis * time.Millisecond)
	keepaliveMsg := pingMsg

	var quit <-chan time.Time

	gotPong := false
	sentPong := false

	for {
		select {
		case message, ok := <-messageChan:
			if !ok {
				return
			}

			switch {
			case string(message) == pingMsg:
				keepaliveMsg = pongMsg

			case string(message) == pongMsg:
				if !gotPong {
					log.Println("Received pong message.")
				}

				// One client may skip sending ping if it receives
				// a ping message before knowning the peer address.
				keepaliveMsg = pongMsg

				gotPong = true

			case stun.IsMessage(message):
				m := new(stun.Message)
				m.Raw = message
				decErr := m.Decode()
				if decErr != nil {
					log.Println("decode:", decErr)

					break
				}
				var xorAddr stun.XORMappedAddress
				if getErr := xorAddr.GetFrom(m); getErr != nil {
					log.Println("getFrom:", getErr)

					break
				}

				if publicAddr.String() != xorAddr.String() {
					log.Printf("My public address: %s\n", xorAddr)
					publicAddr = xorAddr

					peerAddrChan = getPeerAddr()
				}

			default:
				log.Panicln("unknown message", string(message))
			}

		case peerStr := <-peerAddrChan:
			peerAddr, err = net.ResolveUDPAddr(udp, peerStr)
			if err != nil {
				log.Panicln("resolve peeraddr:", err)
			}

		case <-keepalive:
			// Keep NAT binding alive using STUN server or the peer once it's known
			if peerAddr == nil {
				err = sendBindingRequest(conn, srvAddr)
			} else {
				err = sendStr(keepaliveMsg, conn, peerAddr)
				if keepaliveMsg == pongMsg {
					sentPong = true
				}
			}

			if err != nil {
				log.Panicln("keepalive:", err)
			}

		case <-quit:
			_ = conn.Close()
		}

		if quit == nil && gotPong && sentPong {
			log.Println("Success! Quitting in two seconds.")
			quit = time.After(2 * time.Second)
		}
	}
}

func getPeerAddr() <-chan string { _ = "STUB: not implemented"; return nil }

func listen(conn *net.UDPConn) <-chan []byte { _ = "STUB: not implemented"; return nil }

func sendBindingRequest(conn *net.UDPConn, addr *net.UDPAddr) error {
	_ = "STUB: not implemented"
	return nil
}

func send(msg []byte, conn *net.UDPConn, addr *net.UDPAddr) error {
	_ = "STUB: not implemented"
	return nil
}

func sendStr(msg string, conn *net.UDPConn, addr *net.UDPAddr) error {
	_ = "STUB: not implemented"
	return nil
}
