//js:package keychain

package keychain

import "context"

//go:generate go run github.com/grexie/isolates/codegen

type Import interface {}

var _ Keychain = &KeychainBase{}

type Keychain interface {
	Create(ctx context.Context, service string) (Credential, error)
	Get(ctx context.Context, service string) (Credential, error)
	Has(ctx context.Context, service string) (bool, error)
	List(ctx context.Context) ([]Credential, error)
}

type Credential interface {
	Service(ctx context.Context) (string, error)
	Data(ctx context.Context) ([]byte, error)
	SetData(ctx context.Context, bytes []byte) error
}

