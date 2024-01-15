//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLIntersectionFunctionTableDescriptor : objc.NSObject
*/

type MTLIntersectionFunctionTableDescriptor struct {
  *objc.NSObject

}

func (r *MTLIntersectionFunctionTableDescriptor) SetFunctionCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIntersectionFunctionTableDescriptor) IntersectionFunctionTableDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIntersectionFunctionTableDescriptor) FunctionCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

