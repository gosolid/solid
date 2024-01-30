//js:package keychain

package keychain

import (
	"context"
	"fmt"

	"github.com/grexie/isolates"
	keychain "github.com/keybase/go-keychain"
)

//go:generate go run github.com/grexie/isolates/codegen

//js:class Keychain
type KeychainBase struct {
}

//js:class Credential
type CredentialBase struct {
	service string
}

//js:constructor Credential
func newCredential(in isolates.FunctionArgs) (*CredentialBase, error) {
	return &CredentialBase{}, nil
}

//js:constructor Keychain
func newKeychain(in isolates.FunctionArgs) (*KeychainBase, error) {
	return &KeychainBase{}, nil
}

//js:method
//js:export-instance default
func newLibrary(in isolates.FunctionArgs) (*KeychainBase, error) {
	return &KeychainBase{}, nil
} 

//js:method
func (k *KeychainBase) Create(ctx context.Context, service string) (Credential, error) {
	credential := &CredentialBase{
		service: service,
	}
	return credential, nil
}

//js:method
func (k *KeychainBase) Get(ctx context.Context, service string) (Credential, error) {
	credential := &CredentialBase{
		service: service,
	}
	return credential, nil
}

//js:method
func (k *KeychainBase) Has(ctx context.Context, service string) (bool, error) {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(service)
	item.SetAccount(service)
	item.SetLabel(service)
	item.SetReturnAttributes(true)

	if items, err := keychain.QueryItem(item); err != nil {
		return false, err
	} else if len(items) == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

//js:method
func (k *KeychainBase) List(ctx context.Context) ([]Credential, error) {
	return nil, fmt.Errorf("not implemented")
}

//js:get
func (c *CredentialBase) Data(ctx context.Context) ([]byte, error) {
	item := keychain.NewItem()
	item.SetReturnData(true)
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(c.service)
	item.SetAccount(c.service)
	item.SetLabel(c.service)

	if items, err := keychain.QueryItem(item); err != nil {
		return nil, err
	} else if len(items) == 0 {
		return nil, nil
	} else {
		return items[0].Data, nil
	}
}

//js:get
func (c *CredentialBase) Service(ctx context.Context) (string, error) {
	return c.service, nil
}

//js:set data
func (c *CredentialBase) SetData(ctx context.Context, bytes []byte) error {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(c.service)
	item.SetAccount(c.service)
	item.SetLabel(c.service)
	item.SetReturnAttributes(true)

	updateItem := keychain.NewItem()
	updateItem.SetSecClass(keychain.SecClassGenericPassword)
	updateItem.SetService(c.service)
	updateItem.SetAccount(c.service)
	updateItem.SetLabel(c.service)
	updateItem.SetAccessible(keychain.AccessibleAlways)
	updateItem.SetData(bytes)

	if items, err := keychain.QueryItem(item); err != nil {
		return err
	} else if len(items) == 0 {
		return keychain.AddItem(updateItem)
	} else {
		return keychain.UpdateItem(item, updateItem)
	}

}