// Package ripemd160 implements the RIPEMD-160 hash algorithm.
package ripemd160

import (
	"crypto"
	"github.com/ricardobranco777/go-openssl/digest"
	"hash"
)

func init() {
	crypto.RegisterHash(crypto.RIPEMD160, New)
}

const (
	// The block size of the hash algorithm in bytes.
	BlockSize = 64
	// The size of the checksum in bytes.
	Size = 20
)

// New returns a new hash.Hash computing the checksum.
func New() hash.Hash {
	return digest.New("ripemd160")
}

// Sum returns the checksum of the data.
func Sum(data []byte) (md [Size]byte) {
	h := New()
	h.Write(data)
	copy(md[:], h.Sum(nil))
	return
}
