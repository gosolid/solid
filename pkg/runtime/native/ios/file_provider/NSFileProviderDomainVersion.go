//js:package native/ios/file-provider
package file_provider

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface FileProvider.NSFileProviderDomainVersion : objc.NSObject
*/

type NSFileProviderDomainVersion struct {
  *objc.NSObject

}

func (r *NSFileProviderDomainVersion) Next() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderDomainVersion) Compare() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

