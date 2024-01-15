//js:package native/ios/metal
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

func (r *MTLArrayType) ElementType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArrayType) ArrayLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArrayType) Stride() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArrayType) ArgumentIndexStride() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArrayType) ElementStructType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArrayType) ElementArrayType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArrayType) ElementTextureReferenceType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArrayType) ElementPointerType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

