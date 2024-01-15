//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLVisibleFunctionTableDescriptor : objc.NSObject
*/

type MTLVisibleFunctionTableDescriptor struct {
  *objc.NSObject

}

func (r *MTLVisibleFunctionTableDescriptor) VisibleFunctionTableDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVisibleFunctionTableDescriptor) FunctionCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVisibleFunctionTableDescriptor) SetFunctionCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

