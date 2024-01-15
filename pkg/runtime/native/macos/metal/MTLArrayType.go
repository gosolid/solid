//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Metal.MTLArrayType : Metal.MTLType
*/

type MTLArrayType struct {
  *MTLType

}

func (r *MTLArrayType) ArrayLength() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLArrayType) Stride() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLArrayType) ArgumentIndexStride() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLArrayType) ElementStructType() (*MTLStructType, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArrayType) ElementArrayType() (*MTLArrayType, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArrayType) ElementTextureReferenceType() (*MTLTextureReferenceType, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArrayType) ElementPointerType() (*MTLPointerType, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArrayType) ElementType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

