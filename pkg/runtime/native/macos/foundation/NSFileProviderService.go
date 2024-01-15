//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSFileProviderService : objc.NSObject
*/

type NSFileProviderService struct {
  *objc.NSObject

}

func (r *NSFileProviderService) GetFileProviderConnectionWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFileProviderService) Name() (**NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

