// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package stun

import (
	"errors"
	"io"

	"github.com/pion/logging"
)

const (
	// magicCookie is fixed value that aids in distinguishing STUN packets
	// from packets of other protocols when STUN is multiplexed with those
	// other protocols on the same Port.
	//
	// The magic cookie field MUST contain the fixed value 0x2112A442 in
	// network byte order.
	//
	// Defined in "STUN Message Structure", section 6.
	magicCookie         = 0x2112A442
	attributeHeaderSize = 4
	messageHeaderSize   = 20

	// TransactionIDSize is length of transaction id array (in bytes).
	TransactionIDSize = 12 // 96 bit
)

// NewTransactionID returns new random transaction ID using crypto/rand
// as source.
func NewTransactionID() (b [TransactionIDSize]byte) { _ = "STUB: not implemented"; return nil }

// IsMessage returns true if b looks like STUN message.
// Useful for multiplexing. IsMessage does not guarantee
// that decoding will be successful.
func IsMessage(b []byte) bool { _ = "STUB: not implemented"; return false }

// MessageOption is a function that sets a Message option.
type MessageOption func(*Message)

// New returns *Message with pre-allocated Raw.
func New() *Message { _ = "STUB: not implemented"; return nil }

// NewWithOptions returns *Message with pre-allocated Raw, applying the provided options.
func NewWithOptions(options ...MessageOption) *Message { _ = "STUB: not implemented"; return nil }

// ErrDecodeToNil occurs on Decode(data, nil) call.
var ErrDecodeToNil = errors.New("attempt to decode to nil message")

// Decode decodes Message from data to m, returning error if any.
func Decode(data []byte, m *Message) error { _ = "STUB: not implemented"; return nil }

// Message represents a single STUN packet. It uses aggressive internal
// buffering to enable zero-allocation encoding and decoding,
// so there are some usage constraints:
//
//	Message, its fields, results of m.Get or any attribute a.GetFrom
//	are valid only until Message.Raw is not modified.
type Message struct {
	Type          MessageType
	Length        uint32 // len(Raw) not including header
	TransactionID [TransactionIDSize]byte
	Attributes    Attributes
	Raw           []byte
	logger        logging.LeveledLogger
	strict        bool
}

// withMessageLogger sets the logger for the Message.
func withMessageLogger(logger logging.LeveledLogger) MessageOption {
	_ = "STUB: not implemented"
	return *new(MessageOption)
}

// WithStrict enables stricter RFC 5389 and RFC 8489 enforcement.
func WithStrict(strict bool) MessageOption { _ = "STUB: not implemented"; return *new(MessageOption) }

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (m Message) MarshalBinary() (data []byte, err error) {
	_ = "STUB: not implemented"
	// We can't return m.Raw, allocation is expected by implicit interface
	// contract induced by other implementations.
	return nil, nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (m *Message) UnmarshalBinary(data []byte) error {
	_ = "STUB: not implemented"
	// We can't retain data, copy is expected by interface contract.
	return nil
}

// GobEncode implements the gob.GobEncoder interface.
func (m Message) GobEncode() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// GobDecode implements the gob.GobDecoder interface.
		nil
}

func (m *Message) GobDecode(data []byte) error { _ = "STUB: not implemented"; return nil }

// AddTo sets b.TransactionID to m.TransactionID.
//
// Implements Setter to aid in crafting responses.
func (m *Message) AddTo(b *Message) error { _ = "STUB: not implemented"; return nil }

// NewTransactionID sets m.TransactionID to random value from crypto/rand
// and returns error if any.
func (m *Message) NewTransactionID() error { _ = "STUB: not implemented"; return nil }

func (m *Message) String() string { _ = "STUB: not implemented"; return "" }

// Reset resets Message, attributes and underlying buffer length.
func (m *Message) Reset() { _ = "STUB: not implemented"; return }

// grow ensures that internal buffer has n length.
func (m *Message) grow(n int) { _ = "STUB: not implemented"; return }

// Add appends new attribute to message. Not goroutine-safe.
//
// Value of attribute is copied to internal buffer so
// it is safe to reuse v.
func (m *Message) Add(attrType AttrType, val []byte) {
	_ = "STUB: not implemented"
	// Allocating buffer for TLV (type-length-value).
	// T = t, L = len(v), V = v.
	// m.Raw will look like:
	// [0:20]                               <- message header
	// [20:20+m.Length]                     <- existing message attributes
	// [20+m.Length:20+m.Length+len(v) + 4] <- allocated buffer for new TLV
	// [first:last]                         <- same as previous
	// [0 1|2 3|4    4 + len(v)]            <- mapping for allocated buffer
	//
	//	T   L        V
	return
}

// ~ len(TLV) = len(TL) + len(V)
// first byte number
// last byte number
// growing cap(Raw) to fit TLV
// now len(Raw) = last
//nolint:gosec // G115
// rendering length change

// Sub-slicing internal buffer to simplify encoding.
// slice for TLV
// slice for V

// T
//nolint:gosec // G115
// L
// V

// Encoding attribute TLV to allocated buffer.
// T
// L
// V

// Checking that attribute value needs padding.

// Performing padding.

// setting all padding bytes to zero
// to prevent data leak from previous
// data in next bytesToAdd bytes

// increasing buffer length
//nolint:gosec // G115
// rendering length change

func attrSliceEqual(a, b Attributes) bool { _ = "STUB: not implemented"; return false }

func attrEqual(attrA, attrB Attributes) bool { _ = "STUB: not implemented"; return false }

// Equal returns true if Message msg equals to m.
// Ignores m.Raw.
func (m *Message) Equal(msg *Message) bool { _ = "STUB: not implemented"; return false }

// WriteLength writes m.Length to m.Raw.
func (m *Message) WriteLength() { _ = "STUB: not implemented"; return }

//nolint:gosec // G115

// WriteHeader writes header to underlying buffer. Not goroutine-safe.
func (m *Message) WriteHeader() { _ = "STUB: not implemented"; return }

// early bounds check to guarantee safety of writes below

// magic cookie
// transaction ID

// WriteTransactionID writes m.TransactionID to m.Raw.
func (m *Message) WriteTransactionID() { _ = "STUB: not implemented"; return }

// transaction ID

// WriteAttributes encodes all m.Attributes to m.
func (m *Message) WriteAttributes() { _ = "STUB: not implemented"; return }

// WriteType writes m.Type to m.Raw.
func (m *Message) WriteType() { _ = "STUB: not implemented"; return }

// message type

// SetType sets m.Type and writes it to m.Raw.
func (m *Message) SetType(t MessageType) { _ = "STUB: not implemented"; return }

// Encode re-encodes message into m.Raw.
func (m *Message) Encode() { _ = "STUB: not implemented"; return }

// WriteTo implements WriterTo via calling Write(m.Raw) on w and returning
// call result.
func (m *Message) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadFrom implements ReaderFrom. Reads message from r into m.Raw,
// Decodes it and return error if any. If m.Raw is too small, will return
// ErrUnexpectedEOF, ErrUnexpectedHeaderEOF or *DecodeErr.
//
// Can return *DecodeErr while decoding too.
func (m *Message) ReadFrom(r io.Reader) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// ErrUnexpectedHeaderEOF means that there were not enough bytes in
// m.Raw to read header.
var ErrUnexpectedHeaderEOF = errors.New("unexpected EOF: not enough bytes to read header")

// ErrInvalidType means that the message type is 0 (reserved).
var ErrInvalidType = errors.New("STUN message type 0 is reserved")

// Decode decodes m.Raw into m.
func (m *Message) Decode() error {
	_ = "STUB: not implemented" //nolint:cyclop
	// decoding message header
	return nil
}

// first 2 bytes
// second 2 bytes
// last 4 bytes
// len(m.Raw)

// saving header data

//nolint:gosec // G115

// checking that we have enough bytes to read header

// first 2 bytes
// second 2 bytes

// attribute length
// expected buffer length (with padding)

// slicing again to simplify value read

// checking size

// RFC 8489:
// - after MESSAGE-INTEGRITY, only MESSAGE-INTEGRITY-SHA256 and
//   FINGERPRINT may follow.
// - after MESSAGE-INTEGRITY-SHA256, if MESSAGE-INTEGRITY was absent,
//   only FINGERPRINT may follow.

// Write decodes message and return error if any.
//
// Any error is unrecoverable, but message could be partially decoded.
func (m *Message) Write(tBuf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// CloneTo clones m to b securing any further m mutations.
func (m *Message) CloneTo(b *Message) error { _ = "STUB: not implemented"; return nil }

// MessageClass is 8-bit representation of 2-bit class of STUN Message Class.
type MessageClass byte

// Possible values for message class in STUN Message Type.
const (
	ClassRequest         MessageClass = 0x00 // 0b00
	ClassIndication      MessageClass = 0x01 // 0b01
	ClassSuccessResponse MessageClass = 0x02 // 0b10
	ClassErrorResponse   MessageClass = 0x03 // 0b11
)

// Common STUN message types.
var (
	// Binding request message type.
	BindingRequest = NewType(MethodBinding, ClassRequest) //nolint:gochecknoglobals
	// Binding success response message type.
	BindingSuccess = NewType(MethodBinding, ClassSuccessResponse) //nolint:gochecknoglobals
	// Binding error response message type.
	BindingError = NewType(MethodBinding, ClassErrorResponse) //nolint:gochecknoglobals
)

func (c MessageClass) String() string { _ = "STUB: not implemented"; return "" }

//nolint

// Method is uint16 representation of 12-bit STUN method.
type Method uint16

// Possible methods for STUN Message.
const (
	MethodBinding          Method = 0x001
	MethodAllocate         Method = 0x003
	MethodRefresh          Method = 0x004
	MethodSend             Method = 0x006
	MethodData             Method = 0x007
	MethodCreatePermission Method = 0x008
	MethodChannelBind      Method = 0x009
)

// Methods from RFC 6062.
const (
	MethodConnect           Method = 0x000a
	MethodConnectionBind    Method = 0x000b
	MethodConnectionAttempt Method = 0x000c
)

func methodName() map[Method]string { _ = "STUB: not implemented"; return nil }

// RFC 6062.

func (m Method) String() string { _ = "STUB: not implemented"; return "" }

// Falling back to hex representation.

// MessageType is STUN Message Type Field.
type MessageType struct {
	Method Method       // e.g. binding
	Class  MessageClass // e.g. request
}

// AddTo sets m type to t.
func (t MessageType) AddTo(m *Message) error { _ = "STUB: not implemented"; return nil }

// NewType returns new message type with provided method and class.
func NewType(method Method, class MessageClass) MessageType {
	_ = "STUB: not implemented"
	return *new(MessageType)
}

const (
	methodABits = 0xf   // 0b0000000000001111
	methodBBits = 0x70  // 0b0000000001110000
	methodDBits = 0xf80 // 0b0000111110000000

	methodBShift = 1
	methodDShift = 2

	firstBit  = 0x1
	secondBit = 0x2

	c0Bit = firstBit
	c1Bit = secondBit

	classC0Shift = 4
	classC1Shift = 7
)

// Value returns bit representation of messageType.
func (t MessageType) Value() uint16 {
	_ = "STUB: not implemented"
	//	 0                 1
	//	 2  3  4 5 6 7 8 9 0 1 2 3 4 5
	//	+--+--+-+-+-+-+-+-+-+-+-+-+-+-+
	//	|M |M |M|M|M|C|M|M|M|C|M|M|M|M|
	//	|11|10|9|8|7|1|6|5|4|0|3|2|1|0|
	//	+--+--+-+-+-+-+-+-+-+-+-+-+-+-+
	//
	// Figure 3: Format of STUN Message Type Field
	return 0
}

// Warning: Abandon all hope ye who enter here.
// Splitting M into A(M0-M3), B(M4-M6), D(M7-M11).

// A = M * 0b0000000000001111 (right 4 bits)
// B = M * 0b0000000001110000 (3 bits after A)
// D = M * 0b0000111110000000 (5 bits after B)

// Shifting to add "holes" for C0 (at 4 bit) and C1 (8 bit).

// C0 is zero bit of C, C1 is first bit.
// C0 = C * 0b01, C1 = (C * 0b10) >> 1
// Ct = C0 << 4 + C1 << 8.
// Optimizations: "((C * 0b10) >> 1) << 8" as "(C * 0b10) << 7"
// We need C0 shifted by 4, and C1 by 8 to fit "11" and "7" positions
// (see figure 3).

// ReadValue decodes uint16 into MessageType.
func (t *MessageType) ReadValue(v uint16) {
	_ = "STUB: not implemented"
	// Decoding class.
	// We are taking first bit from v >> 4 and second from v >> 7.
	return
}

// Decoding method.
// A(M0-M3)
// B(M4-M6)
// D(M7-M11)

func (t MessageType) String() string { _ = "STUB: not implemented"; return "" }

// Contains return true if message contain t attribute.
func (m *Message) Contains(t AttrType) bool { _ = "STUB: not implemented"; return false }

type transactionIDValueSetter [TransactionIDSize]byte

// NewTransactionIDSetter returns new Setter that sets message transaction id
// to provided value.
func NewTransactionIDSetter(value [TransactionIDSize]byte) Setter {
	_ = "STUB: not implemented"
	return *new(Setter)
}

func (t transactionIDValueSetter) AddTo(m *Message) error { _ = "STUB: not implemented"; return nil }
