//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSAssertionHandler : objc.NSObject
*/

type NSAssertionHandler struct {
  *objc.NSObject

}

func (r *NSAssertionHandler) HandleFailureInMethod() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAssertionHandler) HandleFailureInFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAssertionHandler) CurrentHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

