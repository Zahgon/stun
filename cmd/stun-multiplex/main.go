// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Command stun-multiplex is example of doing UDP connection multiplexing
// that splits incoming UDP packets to two streams, "STUN Data" and
// "Application Data".
package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pion/stun/v3"
)

func copyAddr(dst *stun.XORMappedAddress, src stun.XORMappedAddress) {
	_ = "STUB: not implemented"
	return
}

func keepAlive(c *stun.Client) {
	_ = "STUB: not implemented"
	// Keep-alive for NAT binding.
	return
}

type message struct {
	text string
	addr net.Addr
}

func demultiplex(conn *net.UDPConn, stunConn io.Writer, messages chan message) {
	_ = "STUB: not implemented"
	return
}

// De-multiplexing incoming packets.

// If buf looks like STUN message, send it to STUN client connection.

// If not, it is application data.

func multiplex(conn *net.UDPConn, stunAddr net.Addr, stunConn io.Reader) {
	_ = "STUB: not implemented"
	// Sending all data from stun client to stun server.
	return
}

var stunServer = flag.String("stun", "stun.l.google.com:19302", "STUN Server to use") //nolint:gochecknoglobals

func main() { //nolint:cyclop
	flag.Parse()

	isServer := flag.Arg(0) == ""

	// Allocating local UDP socket that will be used both for STUN and
	// our application data.
	addr, err := net.ResolveUDPAddr("udp4", "0.0.0.0:0")
	if err != nil {
		log.Panicf("Failed to resolve: %s", err)
	}

	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		log.Panicf("Failed to listen: %s", err)
	}

	// Resolving STUN server address.
	stunAddr, err := net.ResolveUDPAddr("udp4", *stunServer)
	if err != nil {
		log.Panicf("Failed to resolve '%s': %s", *stunServer, err)
	}

	log.Printf("Local address: %s", conn.LocalAddr())
	log.Printf("STUN server address: %s", stunAddr)

	stunL, stunR := net.Pipe()

	client, err := stun.NewClient(stunR)
	if err != nil {
		log.Panicf("Failed to create client: %s", err)
	}

	// Starting multiplexing (writing back STUN messages) with de-multiplexing
	// (passing STUN messages to STUN client and processing application
	// data separately).
	//
	// stunL and stunR are virtual connections, see net.Pipe for reference.
	messages := make(chan message)

	go demultiplex(conn, stunL, messages)
	go multiplex(conn, stunAddr, stunL)

	// Getting our "real" IP address from STUN Server.
	// This will create a NAT binding on your provider/router NAT Server,
	// and the STUN server will return allocated public IP for that binding.
	//
	// This can fail if your NAT Server is strict and will use separate ports
	// for application data and STUN
	var gotAddr stun.XORMappedAddress
	if err = client.Do(stun.MustBuild(stun.TransactionID, stun.BindingRequest), func(res stun.Event) {
		if res.Error != nil {
			log.Panicf("Failed STUN transaction: %s", res.Error)
		}
		var xorAddr stun.XORMappedAddress
		if getErr := xorAddr.GetFrom(res.Message); getErr != nil {
			log.Panicf("Failed to get XOR-MAPPED-ADDRESS: %s", getErr)
		}
		copyAddr(&gotAddr, xorAddr)
	}); err != nil {
		log.Panicf("Failed STUN transaction: %s", err)
	}

	log.Printf("Public address: %s", gotAddr)

	// Keep-alive is needed to keep our NAT port allocated.
	// Any ping-pong will work, but we are just making binding requests.
	// Note that STUN Server is not mandatory for keep alive, application
	// data will keep alive that binding too.
	go keepAlive(client)

	notify := make(chan os.Signal, 1)
	signal.Notify(notify, os.Interrupt, syscall.SIGTERM)
	if isServer {
		//nolint:gosec // G705 -- no xss.
		log.Printf("Acting as server. Use following command to connect: %s %s", os.Args[0], gotAddr)

		for {
			select {
			case m := <-messages:
				if _, err = conn.WriteTo([]byte(m.text), m.addr); err != nil {
					log.Panicf("Failed to write: %s", err)
				}
			case <-notify:
				log.Println("Stopping")

				return
			}
		}
	} else {
		peerAddr, err := net.ResolveUDPAddr("udp4", flag.Arg(0))
		if err != nil {
			log.Panicf("Failed to resolve '%s': %s", flag.Arg(0), err)
		}

		log.Printf("Acting as client. Connecting to %s", peerAddr)

		msg := "Hello peer"

		sendMsg := func() {
			log.Printf("Writing to: %s", peerAddr)
			if _, err = conn.WriteTo([]byte(msg), peerAddr); err != nil {
				log.Panicf("Failed to write: %s", err)
			}
		}

		sendMsg()

		deadline := time.After(time.Second * 10)

		for {
			select {
			case <-deadline:
				log.Fatal("Failed to connect: deadline reached.")

			case <-time.After(time.Second):
				// Retry.
				sendMsg()

			case m := <-messages:
				log.Printf("Got response from %s: %s", m.addr, m.text)

				return

			case <-notify:
				log.Print("Stopping")

				return
			}
		}
	}
}
