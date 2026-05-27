// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package stun

import "errors"

// DecodeErr records an error and place when it is occurred.
//
//nolint:errname
type DecodeErr struct {
	Place   DecodeErrPlace
	Message string
}

// IsInvalidCookie returns true if error means that magic cookie
// value is invalid.
func (e DecodeErr) IsInvalidCookie() bool { _ = "STUB: not implemented"; return false }

// IsPlaceParent reports if error place parent is p.
func (e DecodeErr) IsPlaceParent(p string) bool { _ = "STUB: not implemented"; return false }

// IsPlaceChildren reports if error place children is c.
func (e DecodeErr) IsPlaceChildren(c string) bool { _ = "STUB: not implemented"; return false }

// IsPlace reports if error place is p.
func (e DecodeErr) IsPlace(p DecodeErrPlace) bool { _ = "STUB: not implemented"; return false }

// DecodeErrPlace records a place where error is occurred.
type DecodeErrPlace struct {
	Parent   string
	Children string
}

func (p DecodeErrPlace) String() string { _ = "STUB: not implemented"; return "" }

func (e DecodeErr) Error() string { _ = "STUB: not implemented"; return "" }

func newDecodeErr(parent, children, message string) *DecodeErr {
	_ = "STUB: not implemented"
	return nil
}

func newAttrDecodeErr(children, message string) *DecodeErr { _ = "STUB: not implemented"; return nil }

// ErrAttributeSizeInvalid means that decoded attribute size is invalid.
var ErrAttributeSizeInvalid = errors.New("attribute size is invalid")

// ErrAttributeSizeOverflow means that decoded attribute size is too big.
var ErrAttributeSizeOverflow = errors.New("attribute size overflow")

// errInvalidErrorCode means that ErrorCode can't be encoded in ERROR-CODE.
var errInvalidErrorCode = errors.New("invalid ErrorCode")
