//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSURLCredentialStorage : objc.NSObject
*/

type NSURLCredentialStorage struct {
  *objc.NSObject

}

func (r *NSURLCredentialStorage) SetCredential() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLCredentialStorage) RemoveCredential() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLCredentialStorage) DefaultCredentialForProtectionSpace() (*NSURLCredential, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCredentialStorage) SetDefaultCredential() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLCredentialStorage) SharedCredentialStorage() (*NSURLCredentialStorage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCredentialStorage) AllCredentials() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCredentialStorage) CredentialsForProtectionSpace() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

