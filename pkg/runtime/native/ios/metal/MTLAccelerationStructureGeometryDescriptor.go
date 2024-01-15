//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLAccelerationStructureGeometryDescriptor : objc.NSObject
*/

type MTLAccelerationStructureGeometryDescriptor struct {
  *objc.NSObject

}

func (r *MTLAccelerationStructureGeometryDescriptor) IntersectionFunctionTableOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) PrimitiveDataBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) SetPrimitiveDataElementSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) SetPrimitiveDataBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) PrimitiveDataStride() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) Label() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) SetLabel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) PrimitiveDataBufferOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) SetPrimitiveDataBufferOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) PrimitiveDataElementSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) SetIntersectionFunctionTableOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) SetOpaque() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) SetAllowDuplicateIntersectionFunctionInvocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) Opaque() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) AllowDuplicateIntersectionFunctionInvocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureGeometryDescriptor) SetPrimitiveDataStride() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

