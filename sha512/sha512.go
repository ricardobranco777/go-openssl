// Package sha512 implements the SHA-384, SHA-512, SHA-512/224, and SHA-512/256
// hash algorithms as defined in FIPS 180-4.
package sha512

import (
	"crypto"
	"github.com/ricardobranco777/go-openssl/digest"
	"hash"
)

func init() {
	crypto.RegisterHash(crypto.SHA384, New384)
	crypto.RegisterHash(crypto.SHA512, New)
	crypto.RegisterHash(crypto.SHA512_224, New512_224)
	crypto.RegisterHash(crypto.SHA512_256, New512_256)
}

const (
	// Size is the size, in bytes, of a SHA-512 checksum.
	Size = 64

	// Size224 is the size, in bytes, of a SHA-512/224 checksum.
	Size224 = 28

	// Size256 is the size, in bytes, of a SHA-512/256 checksum.
	Size256 = 32

	// Size384 is the size, in bytes, of a SHA-384 checksum.
	Size384 = 48

	// BlockSize is the block size, in bytes, of the SHA-512/224,
	// SHA-512/256, SHA-384 and SHA-512 hash functions.
	BlockSize = 128
)

// New returns a new hash.Hash computing the SHA-512 checksum.
func New() hash.Hash {
	return digest.New("sha512")
}

// New384 returns a new hash.Hash computing the SHA-384 checksum.
func New384() hash.Hash {
	return digest.New("sha384")
}

// New512_224 returns a new hash.Hash computing the SHA-512/224 checksum.
func New512_224() hash.Hash {
	return digest.New("sha512-224")
}

// New512_256 returns a new hash.Hash computing the SHA-512/256 checksum.
func New512_256() hash.Hash {
	return digest.New("sha512-256")
}

// Sum512 returns the SHA512 checksum of the data.
func Sum512(data []byte) (md [Size]byte) {
	h := New()
	h.Write(data)
	copy(md[:], h.Sum(nil))
	return
}

// Sum384 returns the SHA384 checksum of the data.
func Sum384(data []byte) (md [Size384]byte) {
	h := New384()
	h.Write(data)
	copy(md[:], h.Sum(nil))
	return
}

// Sum512_224 returns the Sum512/224 checksum of the data.
func Sum512_224(data []byte) (md [Size224]byte) {
	h := New512_224()
	h.Write(data)
	copy(md[:], h.Sum(nil))
	return
}

// Sum512_256 returns the Sum512/256 checksum of the data.
func Sum512_256(data []byte) (md [Size256]byte) {
	h := New512_256()
	h.Write(data)
	copy(md[:], h.Sum(nil))
	return
}
