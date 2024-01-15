//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSExtensionContext : objc.NSObject
*/

type NSExtensionContext struct {
  *objc.NSObject

}

func (r *NSExtensionContext) CompleteRequestReturningItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExtensionContext) CancelRequestWithError() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExtensionContext) OpenURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExtensionContext) InputItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

