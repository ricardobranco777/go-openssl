LD_LIBRARY_PATH=/usr/local/ssl/lib

test:
	@go test github.com/ricardobranco777/go-openssl/md4
	@go test github.com/ricardobranco777/go-openssl/md5
	@go test github.com/ricardobranco777/go-openssl/ripemd160
	@go test github.com/ricardobranco777/go-openssl/sha1
	@go test github.com/ricardobranco777/go-openssl/sha256
	@go test github.com/ricardobranco777/go-openssl/sha512
