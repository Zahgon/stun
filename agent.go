// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package stun

import (
	"errors"
	"sync"
	"time"
)

// NoopHandler just discards any event.
func NoopHandler() Handler {
	_ = "STUB: not implemented"
	return *

	// NewAgent initializes and returns new Agent with provided handler.
	// If h is nil, the NoopHandler will be used.
	new(Handler)
}

func NewAgent(h Handler) *Agent { _ = "STUB: not implemented"; return nil }

// Agent is low-level abstraction over transaction list that
// handles concurrency (all calls are goroutine-safe) and
// time outs (via Collect call).
type Agent struct {
	// transactions is map of transactions that are currently
	// in progress. Event handling is done in such way when
	// transaction is unregistered before agentTransaction access,
	// minimizing mux lock and protecting agentTransaction from
	// data races via unexpected concurrent access.
	transactions map[transactionID]agentTransaction
	closed       bool       // all calls are invalid if true
	mux          sync.Mutex // protects transactions and closed
	handler      Handler    // handles transactions
}

// Handler handles state changes of transaction.
//
// Handler is called on transaction state change.
// Usage of e is valid only during call, user must
// copy needed fields explicitly.
type Handler func(e Event)

// Event is passed to Handler describing the transaction event.
// Do not reuse outside Handler.
type Event struct {
	TransactionID [TransactionIDSize]byte
	Message       *Message
	Error         error
}

// agentTransaction represents transaction in progress.
// Concurrent access is invalid.
type agentTransaction struct {
	id       transactionID
	deadline time.Time
}

var (
	// ErrTransactionStopped indicates that transaction was manually stopped.
	ErrTransactionStopped = errors.New("transaction is stopped")
	// ErrTransactionNotExists indicates that agent failed to find transaction.
	ErrTransactionNotExists = errors.New("transaction not exists")
	// ErrTransactionExists indicates that transaction with same id is already
	// registered.
	ErrTransactionExists = errors.New("transaction exists with same id")
)

// StopWithError removes transaction from list and calls handler with
// provided error. Can return ErrTransactionNotExists and ErrAgentClosed.
func (a *Agent) StopWithError(id [TransactionIDSize]byte, err error) error {
	_ = "STUB: not implemented"
	return nil
}

// Stop stops transaction by id with ErrTransactionStopped, blocking
// until handler returns.
func (a *Agent) Stop(id [TransactionIDSize]byte) error { _ = "STUB: not implemented"; return nil }

// ErrAgentClosed indicates that agent is in closed state and is unable
// to handle transactions.
var ErrAgentClosed = errors.New("agent is closed")

// Start registers transaction with provided id and deadline.
// Could return ErrAgentClosed, ErrTransactionExists.
//
// Agent handler is guaranteed to be eventually called.
func (a *Agent) Start(id [TransactionIDSize]byte, deadline time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// agentCollectCap is initial capacity for Agent.Collect slices,
// sufficient to make function zero-alloc in most cases.
const agentCollectCap = 100

// ErrTransactionTimeOut indicates that transaction has reached deadline.
var ErrTransactionTimeOut = errors.New("transaction is timed out")

// Collect terminates all transactions that have deadline before provided
// time, blocking until all handlers will process ErrTransactionTimeOut.
// Will return ErrAgentClosed if agent is already closed.
//
// It is safe to call Collect concurrently but makes no sense.
func (a *Agent) Collect(gcTime time.Time) error { _ = "STUB: not implemented"; return nil }

// Doing nothing if agent is closed.
// All transactions should be already closed
// during Close() call.

// Adding all transactions with deadline before gcTime
// to toCall and toRemove slices.
// No allocs if there are less than agentCollectCap
// timed out transactions.

// Un-registering timed out transactions.

// Calling handler does not require locked mutex,
// reducing lock time.

// Sending ErrTransactionTimeOut to handler for all transactions,
// blocking until last one.

// Process incoming message, synchronously passing it to handler.
func (a *Agent) Process(m *Message) error { _ = "STUB: not implemented"; return nil }

// SetHandler sets agent handler to h.
func (a *Agent) SetHandler(h Handler) error { _ = "STUB: not implemented"; return nil }

// Close terminates all transactions with ErrAgentClosed and renders Agent to
// closed state.
func (a *Agent) Close() error { _ = "STUB: not implemented"; return nil }

type transactionID [TransactionIDSize]byte
