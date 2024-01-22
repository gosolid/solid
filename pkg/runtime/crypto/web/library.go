//js:package crypto/web
package web

import (
	"context"
	"crypto/rand"
	"fmt"

	isolates "github.com/grexie/isolates"
)

//go:generate go run github.com/grexie/isolates/codegen

type Import interface{}

//js:export getRandomValues
//js:method getRandomValues
func GetRandomValues(ctx context.Context, arrayBufferView *isolates.Value) (*isolates.Value, error) {
	if length, err := arrayBufferView.GetByteLength(ctx); err != nil {
		return nil, err
	} else {
		bytes := make([]byte, length)
		rand.Read(bytes)
		if err := arrayBufferView.SetBytes(ctx, bytes); err != nil {
			return nil, err
		} else {
			return arrayBufferView, nil
		}
	}
}

//js:export createHmac
//js:method createHmac
func CreateHmac(ctx context.Context) (*isolates.Value, error) {
	return nil, fmt.Errorf("hmac not implemented")
}
