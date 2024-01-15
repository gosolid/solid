//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLFunctionDescriptor : objc.NSObject
*/

type MTLFunctionDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLFunctionDescriptor) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionDescriptor) SetName() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLFunctionDescriptor) ConstantValues() (*MTLFunctionConstantValues, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionDescriptor) SetConstantValues() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLFunctionDescriptor) SetOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLFunctionDescriptor) SetBinaryArchives() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLFunctionDescriptor) FunctionDescriptor() (*MTLFunctionDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionDescriptor) SpecializedName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionDescriptor) SetSpecializedName() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLFunctionDescriptor) Options() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionDescriptor) BinaryArchives() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

