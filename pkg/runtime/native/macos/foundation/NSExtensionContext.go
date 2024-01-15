//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSExtensionContext : objc.NSObject
*/

type NSExtensionContext struct {
  *objc.NSObject

}

func (r *NSExtensionContext) CompleteRequestReturningItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSExtensionContext) CancelRequestWithError() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSExtensionContext) OpenURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSExtensionContext) InputItems() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

