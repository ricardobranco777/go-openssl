package sha3

import (
	"golang.org/x/crypto/sha3"
	"hash"
	"testing"
)

var buf = make([]byte, 16384)

func benchmarkSize(newf func() hash.Hash, b *testing.B, size int) {
	bench := newf()
	b.SetBytes(int64(size))
	sum := make([]byte, bench.Size())
	for i := 0; i < b.N; i++ {
		bench.Reset()
		bench.Write(buf[:size])
		bench.Sum(sum[:0])
	}
}

func BenchmarkHash8Bytes_OpenSSL(b *testing.B) {
	benchmarkSize(New512, b, 8)
}

func BenchmarkHash8Bytes(b *testing.B) {
	benchmarkSize(sha3.New512, b, 8)
}

func BenchmarkHash1K_OpenSSL(b *testing.B) {
	benchmarkSize(New512, b, 1024)
}

func BenchmarkHash1K(b *testing.B) {
	benchmarkSize(sha3.New512, b, 1024)
}

func BenchmarkHash8K_OpenSSL(b *testing.B) {
	benchmarkSize(New512, b, 8192)
}

func BenchmarkHash8K(b *testing.B) {
	benchmarkSize(sha3.New512, b, 8192)
}

func BenchmarkHash16K_OpenSSL(b *testing.B) {
	benchmarkSize(New512, b, 16384)
}

func BenchmarkHash16K(b *testing.B) {
	benchmarkSize(sha3.New512, b, 16384)
}
