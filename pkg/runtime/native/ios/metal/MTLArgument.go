//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Metal.MTLArgument : objc.NSObject
*/

type MTLArgument struct {
  *objc.NSObject

}

func (r *MTLArgument) Type() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) BufferAlignment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) TextureType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) IsDepthTexture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) Index() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) IsActive() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) BufferDataSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) ThreadgroupMemoryDataSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) TextureDataType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) Access() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) BufferDataType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) BufferStructType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) BufferPointerType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) ThreadgroupMemoryAlignment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) ArrayLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

