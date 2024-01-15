//js:package native/ios/file-provider
package file_provider

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface FileProvider.NSFileProviderRequest : objc.NSObject
*/

type NSFileProviderRequest struct {
  *objc.NSObject

}

func (r *NSFileProviderRequest) IsSystemRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderRequest) IsFileViewerRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderRequest) RequestingExecutable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderRequest) DomainVersion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

