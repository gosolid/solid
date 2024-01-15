//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface Metal.MTLArgumentDescriptor : objc.NSObject
*/

type MTLArgumentDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLArgumentDescriptor) SetIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) SetArrayLength() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) SetTextureType() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) SetConstantBlockAlignment() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) DataType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) SetDataType() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) Access() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) ArgumentDescriptor() (*MTLArgumentDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) ArrayLength() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) TextureType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) Index() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) SetAccess() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) ConstantBlockAlignment() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

