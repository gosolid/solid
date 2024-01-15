//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLFunctionDescriptor : objc.NSObject
*/

type MTLFunctionDescriptor struct {
  *objc.NSObject

}

func (r *MTLFunctionDescriptor) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionDescriptor) SpecializedName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionDescriptor) SetSpecializedName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionDescriptor) ConstantValues() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionDescriptor) BinaryArchives() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionDescriptor) FunctionDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionDescriptor) SetConstantValues() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionDescriptor) Options() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionDescriptor) SetOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionDescriptor) SetBinaryArchives() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionDescriptor) SetName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

