// Package sha1 implements the SHA1 hash algorithm as defined in RFC 3174.
package sha1

import (
	"crypto"
	"github.com/ricardobranco777/go-openssl/digest"
	"hash"
)

func init() {
	crypto.RegisterHash(crypto.SHA1, New)
}

const (
	// The blocksize of SHA1 in bytes.
	BlockSize = 64
	// The size of a SHA1 checksum in bytes.
	Size = 20
)

// New returns a new hash.Hash computing the SHA1 checksum.
func New() hash.Hash {
	return digest.New("sha1")
}

// Sum returns the SHA1 checksum of the data.
func Sum(data []byte) (md [Size]byte) {
	h := New()
	h.Write(data)
	copy(md[:], h.Sum(nil))
	return
}
