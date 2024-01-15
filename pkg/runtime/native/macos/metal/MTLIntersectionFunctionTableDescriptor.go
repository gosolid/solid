//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLIntersectionFunctionTableDescriptor : objc.NSObject
*/

type MTLIntersectionFunctionTableDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLIntersectionFunctionTableDescriptor) IntersectionFunctionTableDescriptor() (*MTLIntersectionFunctionTableDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIntersectionFunctionTableDescriptor) FunctionCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLIntersectionFunctionTableDescriptor) SetFunctionCount() error {
  return fmt.Errorf("unimplemented")
}

