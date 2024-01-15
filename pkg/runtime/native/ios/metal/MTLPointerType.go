//js:package native/ios/metal
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

func (r *MTLPointerType) Access() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPointerType) Alignment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPointerType) DataSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPointerType) ElementIsArgumentBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPointerType) ElementStructType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPointerType) ElementArrayType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPointerType) ElementType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

