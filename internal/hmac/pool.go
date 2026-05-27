// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package hmac

import (
	"crypto/sha1" //nolint:gosec
	"crypto/sha256"
	"hash"
	"sync"
)

func (h *hmac) resetTo(key []byte) { _ = "STUB: not implemented"; return }

// Reset size and zero of ipad and opad.

// If key is too big, hash it.
//nolint:errcheck,gosec

//nolint:errcheck,gosec

var hmacSHA1Pool = &sync.Pool{ //nolint:gochecknoglobals
	New: func() any {
		h := New(sha1.New, make([]byte, sha1.BlockSize))

		return h
	},
}

// AcquireSHA1 returns new HMAC from pool.
func AcquireSHA1(key []byte) hash.Hash { _ = "STUB: not implemented"; return *new(hash.Hash) }

//nolint:forcetypeassert

// PutSHA1 puts h to pool.
func PutSHA1(h hash.Hash) {
	_ = "STUB: not implemented"
	//nolint:forcetypeassert
	return
}

var hmacSHA256Pool = &sync.Pool{ //nolint:gochecknoglobals
	New: func() any {
		h := New(sha256.New, make([]byte, sha256.BlockSize))

		return h
	},
}

// AcquireSHA256 returns new HMAC from SHA256 pool.
func AcquireSHA256(key []byte) hash.Hash { _ = "STUB: not implemented"; return *new(hash.Hash) }

//nolint:forcetypeassert

// PutSHA256 puts h to SHA256 pool.
func PutSHA256(h hash.Hash) {
	_ = "STUB: not implemented"
	//nolint:forcetypeassert
	return
}

// assertHMACSize panics if h.size != size or h.blocksize != blocksize.
//
// Put and Acquire functions are internal functions to project, so
// checking it via such assert is optimal.
func assertHMACSize(h *hmac, size, blocksize int) {
	_ = "STUB: not implemented" //nolint:unparam
	return
}

//nolint
