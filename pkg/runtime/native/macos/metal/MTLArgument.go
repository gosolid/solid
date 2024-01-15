//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface Metal.MTLArgument : objc.NSObject
*/

type MTLArgument struct {
  *objc.NSObject

}

func (r *MTLArgument) IsActive() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) BufferPointerType() (*MTLPointerType, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) Index() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) ThreadgroupMemoryAlignment() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) ThreadgroupMemoryDataSize() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) TextureDataType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) BufferDataType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) Access() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) BufferAlignment() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) BufferDataSize() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) TextureType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) ArrayLength() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) Type() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) BufferStructType() (*MTLStructType, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) IsDepthTexture() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLArgument) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

