//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Metal.MTLPointerType : Metal.MTLType
*/

type MTLPointerType struct {
  *MTLType

}

func (r *MTLPointerType) ElementStructType() (*MTLStructType, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPointerType) ElementArrayType() (*MTLArrayType, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPointerType) ElementType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLPointerType) Access() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLPointerType) Alignment() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLPointerType) DataSize() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLPointerType) ElementIsArgumentBuffer() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

