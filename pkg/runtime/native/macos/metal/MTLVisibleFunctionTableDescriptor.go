//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLVisibleFunctionTableDescriptor : objc.NSObject
*/

type MTLVisibleFunctionTableDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLVisibleFunctionTableDescriptor) VisibleFunctionTableDescriptor() (*MTLVisibleFunctionTableDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVisibleFunctionTableDescriptor) FunctionCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLVisibleFunctionTableDescriptor) SetFunctionCount() error {
  return fmt.Errorf("unimplemented")
}

