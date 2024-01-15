//js:package native/ios/file-provider
package file_provider

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface FileProvider.NSFileProviderManager : objc.NSObject
*/

type NSFileProviderManager struct {
  *objc.NSObject

}

func (r *NSFileProviderManager) RegisterURLSessionTask() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderManager) WritePlaceholderAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderManager) SignalErrorResolved() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderManager) GlobalProgressForKind() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderManager) DefaultManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderManager) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderManager) RemoveDomain() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderManager) RemoveAllDomainsWithCompletionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderManager) ProviderIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderManager) SignalEnumeratorForContainerItemIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderManager) TemporaryDirectoryURLWithError() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderManager) PlaceholderURLForURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderManager) AddDomain() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderManager) GetDomainsWithCompletionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderManager) ManagerForDomain() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderManager) GetUserVisibleURLForItemIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderManager) GetIdentifierForUserVisibleFileAtURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderManager) DocumentStorageURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

