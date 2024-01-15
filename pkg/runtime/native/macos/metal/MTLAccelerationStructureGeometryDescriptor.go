//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLAccelerationStructureGeometryDescriptor : objc.NSObject
*/

type MTLAccelerationStructureGeometryDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLAccelerationStructureGeometryDescriptor) SetPrimitiveDataBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) SetOpaque() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) SetAllowDuplicateIntersectionFunctionInvocation() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) SetIntersectionFunctionTableOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) SetPrimitiveDataBufferOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) SetPrimitiveDataStride() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) PrimitiveDataBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) PrimitiveDataStride() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) SetPrimitiveDataElementSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) IntersectionFunctionTableOffset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) Opaque() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) AllowDuplicateIntersectionFunctionInvocation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) PrimitiveDataBufferOffset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) PrimitiveDataElementSize() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

