// Package sha256 implements the SHA224 and SHA256 hash algorithms as defined in FIPS 180-4.
package sha256

import (
	"crypto"
	"github.com/ricardobranco777/go-openssl/digest"
	"hash"
)

func init() {
	crypto.RegisterHash(crypto.SHA224, New224)
	crypto.RegisterHash(crypto.SHA256, New)
}

const (
	// The blocksize of SHA256 and SHA224 in bytes.
	BlockSize = 64
	// The size of a SHA256 checksum in bytes.
	Size = 32
	// The size of a SHA224 checksum in bytes.
	Size224 = 28
)

// New returns a new hash.Hash computing the SHA-256 checksum.
func New() hash.Hash {
	return digest.New("sha256")
}

// New224 returns a new hash.Hash computing the SHA224 checksum.
func New224() hash.Hash {
	return digest.New("sha224")
}

// Sum256 returns the SHA256 checksum of the data.
func Sum256(data []byte) (md [Size]byte) {
	h := New()
	h.Write(data)
	copy(md[:], h.Sum(nil))
	return
}

// Sum224 returns the SHA224 checksum of the data.
func Sum224(data []byte) (md [Size224]byte) {
	h := New224()
	h.Write(data)
	copy(md[:], h.Sum(nil))
	return
}
