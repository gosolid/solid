//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLArgumentDescriptor : objc.NSObject
*/

type MTLArgumentDescriptor struct {
  *objc.NSObject

}

func (r *MTLArgumentDescriptor) TextureType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) ConstantBlockAlignment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) SetIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) SetAccess() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) ArgumentDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) ArrayLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) SetTextureType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) SetConstantBlockAlignment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) Index() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) SetDataType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) SetArrayLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) Access() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgumentDescriptor) DataType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

