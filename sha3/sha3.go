package sha3

import (
	"crypto"
	"github.com/ricardobranco777/go-openssl/digest"
	"hash"
)

func init() {
	crypto.RegisterHash(crypto.SHA3_224, New224)
	crypto.RegisterHash(crypto.SHA3_256, New256)
	crypto.RegisterHash(crypto.SHA3_384, New384)
	crypto.RegisterHash(crypto.SHA3_512, New512)
}

const (
	// Size224 is the size, in bytes, of a SHA-512/224 checksum.
	Size224 = 28

	// Size256 is the size, in bytes, of a SHA-512/256 checksum.
	Size256 = 32

	// Size384 is the size, in bytes, of a SHA-384 checksum.
	Size384 = 48

	// Size is the size, in bytes, of a SHA-512 checksum.
	Size512 = 64
)

// New224 creates a new SHA3-224 hash.
// Its generic security strength is 224 bits against preimage attacks,
// and 112 bits against collision attacks.
func New224() hash.Hash {
	return digest.New("sha3-224")
}

// New256 creates a new SHA3-256 hash.
// Its generic security strength is 256 bits against preimage attacks,
// and 128 bits against collision attacks.
func New256() hash.Hash {
	return digest.New("sha3-256")
}

// New384 creates a new SHA3-384 hash.
// Its generic security strength is 384 bits against preimage attacks,
// and 192 bits against collision attacks.
func New384() hash.Hash {
	return digest.New("sha3-384")
}

// New512 creates a new SHA3-512 hash.
// Its generic security strength is 512 bits against preimage attacks,
// and 256 bits against collision attacks.
func New512() hash.Hash {
	return digest.New("sha3-512")
}

// Sum512 returns the SHA512 checksum of the data.
func Sum512(data []byte) (md [Size512]byte) {
	h := New512()
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
	h := New224()
	h.Write(data)
	copy(md[:], h.Sum(nil))
	return
}

// Sum512_256 returns the Sum512/256 checksum of the data.
func Sum512_256(data []byte) (md [Size256]byte) {
	h := New256()
	h.Write(data)
	copy(md[:], h.Sum(nil))
	return
}
