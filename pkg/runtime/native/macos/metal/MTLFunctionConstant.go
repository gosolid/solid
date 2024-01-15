//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface Metal.MTLFunctionConstant : objc.NSObject
*/

type MTLFunctionConstant struct {
  *objc.NSObject

}

func (r *MTLFunctionConstant) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionConstant) Type() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionConstant) Index() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionConstant) Required() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

