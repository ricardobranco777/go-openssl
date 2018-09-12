SSL_LIB_DIR = $(shell ls -d /usr/local/ssl/lib64)
SSL_LIB_DIR ?= /usr/local/ssl/lib

test:
	@LD_LIBRARY_PATH=${SSL_LIB_DIR} go test github.com/ricardobranco777/go-openssl/md4
	@LD_LIBRARY_PATH=${SSL_LIB_DIR} go test github.com/ricardobranco777/go-openssl/md5
	@LD_LIBRARY_PATH=${SSL_LIB_DIR} go test github.com/ricardobranco777/go-openssl/ripemd160
	@LD_LIBRARY_PATH=${SSL_LIB_DIR} go test github.com/ricardobranco777/go-openssl/sha1
	@LD_LIBRARY_PATH=${SSL_LIB_DIR} go test github.com/ricardobranco777/go-openssl/sha256
	@LD_LIBRARY_PATH=${SSL_LIB_DIR} go test github.com/ricardobranco777/go-openssl/sha512

bench:
	@LD_LIBRARY_PATH=${SSL_LIB_DIR} go test github.com/ricardobranco777/go-openssl/md4 -bench=.
	@LD_LIBRARY_PATH=${SSL_LIB_DIR} go test github.com/ricardobranco777/go-openssl/md5 -bench=.
	@LD_LIBRARY_PATH=${SSL_LIB_DIR} go test github.com/ricardobranco777/go-openssl/ripemd160 -bench=.
	@LD_LIBRARY_PATH=${SSL_LIB_DIR} go test github.com/ricardobranco777/go-openssl/sha1 -bench=.
	@LD_LIBRARY_PATH=${SSL_LIB_DIR} go test github.com/ricardobranco777/go-openssl/sha256 -bench=.
	@LD_LIBRARY_PATH=${SSL_LIB_DIR} go test github.com/ricardobranco777/go-openssl/sha512 -bench=.
	@LD_LIBRARY_PATH=${SSL_LIB_DIR} go test github.com/ricardobranco777/go-openssl/sha3 -bench=.
