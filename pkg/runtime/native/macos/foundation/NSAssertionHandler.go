//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Foundation.NSAssertionHandler : objc.NSObject
*/

type NSAssertionHandler struct {
  *objc.NSObject

}

func (r *NSAssertionHandler) CurrentHandler() (*NSAssertionHandler, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAssertionHandler) HandleFailureInMethod() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAssertionHandler) HandleFailureInFunction() error {
  return fmt.Errorf("unimplemented")
}

