//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSScriptCoercionHandler : objc.NSObject
*/

type NSScriptCoercionHandler struct {
  *objc.NSObject

}

func (r *NSScriptCoercionHandler) SharedCoercionHandler() (*NSScriptCoercionHandler, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCoercionHandler) CoerceValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptCoercionHandler) RegisterCoercer() error {
  return fmt.Errorf("unimplemented")
}

