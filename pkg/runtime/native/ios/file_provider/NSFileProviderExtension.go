//js:package native/ios/file-provider
package file_provider

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface FileProvider.NSFileProviderExtension : objc.NSObject
*/

type NSFileProviderExtension struct {
  *objc.NSObject

}

func (r *NSFileProviderExtension) ItemForIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderExtension) URLForItemWithPersistentIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderExtension) PersistentIdentifierForItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderExtension) ProvidePlaceholderAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderExtension) StartProvidingItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderExtension) StopProvidingItemAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderExtension) ItemChangedAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

