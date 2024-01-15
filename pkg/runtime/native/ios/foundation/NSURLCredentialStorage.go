//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Foundation.NSURLCredentialStorage : objc.NSObject
*/

type NSURLCredentialStorage struct {
  *objc.NSObject

}

func (r *NSURLCredentialStorage) AllCredentials() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCredentialStorage) CredentialsForProtectionSpace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCredentialStorage) SetCredential() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCredentialStorage) RemoveCredential() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCredentialStorage) DefaultCredentialForProtectionSpace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCredentialStorage) SetDefaultCredential() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCredentialStorage) SharedCredentialStorage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

