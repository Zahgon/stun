// SPDX-FileCopyrightText: 2009 The Go Authors. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

/*
Package hmac implements the Keyed-Hash Message Authentication Code (HMAC) as
defined in U.S. Federal Information Processing Standards Publication 198.
An HMAC is a cryptographic hash that uses a key to sign a message.
The receiver verifies the hash by recomputing it using the same key.

Receivers should be careful to use Equal to compare MACs in order to avoid
timing side-channels:

	// ValidMAC reports whether messageMAC is a valid HMAC tag for message.
	func ValidMAC(message, messageMAC, key []byte) bool {
		mac := hmac.New(sha256.New, key)
		mac.Write(message)
		expectedMAC := mac.Sum(nil)
		return hmac.Equal(messageMAC, expectedMAC)
	}
*/
package hmac

import (
	"hash"
)

// FIPS 198-1:
// https://csrc.nist.gov/publications/fips/fips198-1/FIPS-198-1_final.pdf

// key is zero padded to the block size of the hash function
// ipad = 0x36 byte repeated for key length
// opad = 0x5c byte repeated for key length
// hmac = H([key ^ opad] H([key ^ ipad] text))

// Marshalable is the combination of encoding.BinaryMarshaler and
// encoding.BinaryUnmarshaler. Their method definitions are repeated here to
// avoid a dependency on the encoding package.
type marshalable interface {
	MarshalBinary() ([]byte, error)
	UnmarshalBinary([]byte) error
}

type hmac struct {
	opad, ipad   []byte
	outer, inner hash.Hash

	// If marshaled is true, then opad and ipad do not contain a padded
	// copy of the key, but rather the marshaled state of outer/inner after
	// opad/ipad has been fed into it.
	marshaled bool
}

func (h *hmac) Sum(in []byte) []byte { _ = "STUB: not implemented"; return nil }

//nolint:forcetypeassert
//nolint

//nolint:errcheck,gosec

//nolint:errcheck,gosec

func (h *hmac) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (h *hmac) Size() int      { _ = "STUB: not implemented"; return 0 }
func (h *hmac) BlockSize() int { _ = "STUB: not implemented"; return 0 }

func (h *hmac) Reset() { _ = "STUB: not implemented"; return }

//nolint:forcetypeassert
//nolint

//nolint:errcheck,gosec

// If the underlying hash is marshalable, we can save some time by
// saving a copy of the hash state now, and restoring it on future
// calls to Reset and Sum instead of writing ipad/opad every time.
//
// If either hash is unmarshalable for whatever reason,
// it's safe to bail out here.

//nolint:errcheck,gosec

// Marshaling succeeded; save the marshaled state for later

// New returns a new HMAC hash using the given hash.Hash type and key.
// Note that unlike other hash implementations in the standard library,
// the returned Hash does not implement encoding.BinaryMarshaler
// or encoding.BinaryUnmarshaler.
func New(h func() hash.Hash, key []byte) hash.Hash {
	_ = "STUB: not implemented"
	return *new(hash.Hash)
}

// If key is too big, hash it.
//nolint:errcheck,gosec

//nolint:errcheck,gosec

// Equal compares two MACs for equality without leaking timing information.
func Equal(mac1, mac2 []byte) bool {
	_ = "STUB: not implemented"
	// We don't have to be constant time if the lengths of the MACs are
	// different as that suggests that a completely different hash function
	// was used.
	return false
}
