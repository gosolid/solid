//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Metal.MTLStageInputOutputDescriptor : objc.NSObject
*/

type MTLStageInputOutputDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLStageInputOutputDescriptor) Attributes() (*MTLAttributeDescriptorArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStageInputOutputDescriptor) IndexType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLStageInputOutputDescriptor) SetIndexType() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLStageInputOutputDescriptor) IndexBufferIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLStageInputOutputDescriptor) SetIndexBufferIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLStageInputOutputDescriptor) StageInputOutputDescriptor() (*MTLStageInputOutputDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStageInputOutputDescriptor) Reset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLStageInputOutputDescriptor) Layouts() (*MTLBufferLayoutDescriptorArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

