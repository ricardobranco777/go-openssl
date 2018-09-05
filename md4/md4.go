// Package md4 implements the MD4 hash algorithm as defined in RFC 1320.
package md4

import (
	"crypto"
	"github.com/ricardobranco777/go-openssl/digest"
	"hash"
)

func init() {
	crypto.RegisterHash(crypto.MD4, New)
}

const (
	// The blocksize of MD4 in bytes.
	BlockSize = 64
	// The size of an MD4 checksum in bytes.
	Size = 16
)

// New returns a new hash.Hash computing the MD4 checksum.
func New() hash.Hash {
	return digest.New("md4")
}

// Sum returns the MD4 checksum of the data.
func Sum(data []byte) (md [Size]byte) {
	h := New()
	h.Write(data)
	copy(md[:], h.Sum(nil))
	return
}
