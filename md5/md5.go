// Package md5 implements the MD5 hash algorithm as defined in RFC 1321.
package md5

import (
	"crypto"
	"github.com/ricardobranco777/go-openssl/digest"
	"hash"
)

func init() {
	crypto.RegisterHash(crypto.MD5, New)
}

const (
	// The blocksize of MD5 in bytes.
	BlockSize = 64
	// The size of an MD5 checksum in bytes.
	Size = 16
)

// New returns a new hash.Hash computing the MD5 checksum.
func New() hash.Hash {
	return digest.New("md5")
}

// Sum returns the MD5 checksum of the data.
func Sum(data []byte) (md [Size]byte) {
	h := New()
	h.Write(data)
	copy(md[:], h.Sum(nil))
	return
}
