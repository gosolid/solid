//js:package native/ios/file-provider
package file_provider

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface FileProvider.NSFileProviderItemVersion : objc.NSObject
*/

type NSFileProviderItemVersion struct {
  *objc.NSObject

}

func (r *NSFileProviderItemVersion) MetadataVersion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderItemVersion) InitWithContentVersion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderItemVersion) BeforeFirstSyncComponent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderItemVersion) ContentVersion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

