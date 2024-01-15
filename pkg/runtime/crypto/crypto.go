//js:package native:@grexie/workers/crypto
package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"hash/crc32"

	"github.com/gosolid/solid/pkg/runtime/buffer"
)

var _ buffer.Buffer

//go:generate go run github.com/grexie/isolates/codegen

type Hash struct {
	algorithm string
	hash      hash.Hash
}

//js:constructor Hash
//js:export Hash
func NewHash(algorithm string) (*Hash, error) {
	var hash hash.Hash
	switch algorithm {
	case "crc32-ieee":
	case "crc32":
		hash = crc32.NewIEEE()
		break
	case "sha":
	case "sha-1":
	case "sha1":
		hash = sha1.New()
		break
	case "sha-256":
	case "sha256":
		hash = sha256.New()
		break
	case "sha-384":
	case "sha512-384":
	case "sha-512-384":
	case "sha384":
		hash = sha512.New384()
		break
	case "sha-512":
	case "sha512":
		hash = sha512.New()
		break
	case "md5":
		hash = md5.New()
		break
	}
	return &Hash{algorithm, hash}, nil
}

//js:get algorithm
func (h *Hash) Algorithm() string {
	return h.algorithm
}

//js:method update
func (h *Hash) Write(b []byte) (int, error) {
	return h.hash.Write(b)
}

//js:method digest
//js:return buffer.Buffer
func (h *Hash) Sum(b []byte) []byte {
	return h.hash.Sum(b)
}
