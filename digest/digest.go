// Copyright (C) 2016, 2018 Ricardo Branco
//
// MIT License

// +build cgo

// Package digest is a wrapper to OpenSSL EVP_MD
package digest

/*
#cgo CFLAGS: -I/usr/local/ssl/include
#cgo LDFLAGS: -lcrypto -L/usr/local/ssl/lib
#include <openssl/evp.h>
*/
import "C"

import (
	"errors"
	"hash"
	"runtime"
	"unsafe"
)

type digest struct {
	ctx    *C.EVP_MD_CTX
	ctx2   *C.EVP_MD_CTX
	evp_md *C.EVP_MD
	size   int
}

// New returns a new hash.Hash for the specified algorithm
func New(name string) hash.Hash {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	evp_md := C.EVP_get_digestbyname(cname)
	if evp_md == nil {
		return nil
	}
	d := &digest{evp_md: evp_md, size: int(C.EVP_MD_size(evp_md))}
	d.ctx = C.EVP_MD_CTX_new()
	if d.ctx == nil {
		return nil
	}
	d.ctx2 = C.EVP_MD_CTX_new()
	if d.ctx2 == nil {
		return nil
	}
	runtime.SetFinalizer(d, func(d *digest) { C.EVP_MD_CTX_free(d.ctx); C.EVP_MD_CTX_free(d.ctx2) })
	d.Reset()
	return d
}

// Reset resets the Hash to its initial state.
func (d *digest) Reset() {
	C.EVP_DigestInit_ex(d.ctx, d.evp_md, nil)
}

// Write adds more data to the running hash.
func (d *digest) Write(p []byte) (n int, err error) {
	n = len(p)
	if n == 0 {
		return
	}
	if 1 != C.EVP_DigestUpdate(d.ctx, unsafe.Pointer(&p[0]), C.size_t(n)) {
		return 0, errors.New("EVP_DigestUpdate() failed")
	}
	return
}

// Sum appends the current hash to b and returns the resulting slice.
// It does not change the underlying hash state.
func (d *digest) Sum(p []byte) []byte {
	md := make([]byte, d.size)
	C.EVP_MD_CTX_copy_ex(d.ctx2, d.ctx)
	C.EVP_DigestFinal_ex(d.ctx2, (*C.uchar)(unsafe.Pointer(&md[0])), nil)
	return append(p, md...)
}

// Size returns the number of bytes Sum will return.
func (d *digest) Size() int {
	return d.size
}

// BlockSize returns the hash's underlying block size.
func (d *digest) BlockSize() int {
	return int(C.EVP_MD_block_size(d.evp_md))
}
