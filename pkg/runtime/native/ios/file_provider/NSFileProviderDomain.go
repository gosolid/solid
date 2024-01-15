//js:package native/ios/file-provider
package file_provider

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface FileProvider.NSFileProviderDomain : objc.NSObject
*/

type NSFileProviderDomain struct {
  *objc.NSObject

}

func (r *NSFileProviderDomain) SetHidden() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderDomain) BackingStoreIdentity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderDomain) SupportsSyncingTrash() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderDomain) DisplayName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderDomain) PathRelativeToDocumentStorage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderDomain) IsDisconnected() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderDomain) IsReplicated() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderDomain) SetTestingModes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderDomain) InitWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderDomain) IsHidden() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderDomain) TestingModes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderDomain) UserEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderDomain) SetSupportsSyncingTrash() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderDomain) Identifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

