//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Metal.MTLFunctionConstant : objc.NSObject
*/

type MTLFunctionConstant struct {
  *objc.NSObject

}

func (r *MTLFunctionConstant) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionConstant) Type() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionConstant) Index() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionConstant) Required() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

