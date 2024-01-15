//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Foundation.NSFileProviderService : objc.NSObject
*/

type NSFileProviderService struct {
  *objc.NSObject

}

func (r *NSFileProviderService) GetFileProviderConnectionWithCompletionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFileProviderService) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

