//js:package crypto
package crypto

import (
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"hash/crc32"

	"github.com/gosolid/solid/pkg/runtime/buffer"
	isolates "github.com/grexie/isolates"
	"golang.org/x/crypto/ripemd160"
)

var _ buffer.Buffer

//go:generate go run github.com/grexie/isolates/codegen

type Hash struct {
	algorithm string
	hash      hash.Hash
}

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
	case "ripemd160":
		hash = ripemd160.New()
		break
	}

	if hash == nil {
		return nil, fmt.Errorf("invalid hash algorithm: %s", algorithm)
	}

	return &Hash{algorithm, hash}, nil
}

//js:get algorithm
func (h *Hash) Algorithm() string {
	return h.algorithm
}

//js:method update
func (h *Hash) Update(ctx context.Context, b any) (*Hash, error) {
	var bytes []byte
	if bv, err := isolates.For(ctx).Context().CreateWithoutMarshallers(ctx, b); err != nil {
		return nil, err
	} else if bv.IsKind(isolates.KindArrayBuffer) || bv.IsKind(isolates.KindArrayBufferView) {
		if _b, err := bv.Bytes(ctx); err != nil {
			return nil, err
		} else {
			bytes = _b
		}
	} else if bv.IsKind(isolates.KindString) {
		if s, err := bv.StringValue(ctx); err != nil {
			return nil, err
		} else {
			bytes = []byte(s)
		}
	} else {
		return nil, fmt.Errorf("invalid hash update value supplied")
	}

	if _, err := h.hash.Write(bytes); err != nil {
		return nil, err
	} else {
		return h, nil
	}
}

//js:method digest
//js:return buffer.Buffer
func (h *Hash) Sum(encoding *buffer.BufferEncoding) (any, error) {
	digest := h.hash.Sum([]byte{})

	if encoding == nil {
		return digest, nil
	} else if *encoding == buffer.Hex {
		return hex.EncodeToString(digest), nil
	} else if *encoding == buffer.Base64 {
		return base64.RawStdEncoding.EncodeToString(digest), nil
	} else {
		return nil, fmt.Errorf("invalid encoding: %s", encoding)
	}
}

//js:export createHash
//js:method createHash
func CreateHash(ctx context.Context, algorithm string) (*isolates.Value, error) {
	if hash, err := NewHash(algorithm); err != nil {
		return nil, err
	} else if hashv, err := isolates.For(ctx).Context().Create(ctx, hash); err != nil {
		return nil, err
	} else {
		return hashv, nil
	}
}

func NewHmac(algorithm string, key []byte) (*Hash, error) {
	var h hash.Hash

	switch algorithm {
	case "crc32-ieee":
	case "crc32":
		h = hmac.New(func() hash.Hash { return crc32.NewIEEE() }, key)
		break
	case "sha":
	case "sha-1":
	case "sha1":
		h = hmac.New(func() hash.Hash { return sha1.New() }, key)
		break
	case "sha-256":
	case "sha256":
		h = hmac.New(func() hash.Hash { return sha256.New() }, key)
		break
	case "sha-384":
	case "sha512-384":
	case "sha-512-384":
	case "sha384":
		h = hmac.New(func() hash.Hash { return sha512.New384() }, key)
		break
	case "sha-512":
	case "sha512":
		h = hmac.New(func() hash.Hash { return sha512.New() }, key)
		break
	case "md5":
		h = hmac.New(func() hash.Hash { return md5.New() }, key)
		break
	case "ripemd160":
		h = hmac.New(func() hash.Hash { return ripemd160.New() }, key)
		break
	}

	if h == nil {
		return nil, fmt.Errorf("invalid hmac algorithm: %s", algorithm)
	}

	return &Hash{algorithm, h}, nil
}

//js:export createHmac
//js:method createHmac
func CreateHmac(ctx context.Context, algorithm string, key any) (*isolates.Value, error) {
	var keyBytes []byte
	if bv, err := isolates.For(ctx).Context().CreateWithoutMarshallers(ctx, key); err != nil {
		return nil, err
	} else if bv.IsKind(isolates.KindArrayBuffer) || bv.IsKind(isolates.KindArrayBufferView) {
		if _b, err := bv.Bytes(ctx); err != nil {
			return nil, err
		} else {
			keyBytes = _b
		}
	} else if bv.IsKind(isolates.KindString) {
		if s, err := bv.StringValue(ctx); err != nil {
			return nil, err
		} else {
			keyBytes = []byte(s)
		}
	} else {
		return nil, fmt.Errorf("invalid hmac key value supplied")
	}

	if hash, err := NewHmac(algorithm, keyBytes); err != nil {
		return nil, err
	} else if hashv, err := isolates.For(ctx).Context().Create(ctx, hash); err != nil {
		return nil, err
	} else {
		return hashv, nil
	}
}

//js:export randomBytes
//js:method randomBytes
func RandomBytes(ct context.Context) (*isolates.Value, error) {
	return nil, fmt.Errorf("randomBytes unimplemented")
}
