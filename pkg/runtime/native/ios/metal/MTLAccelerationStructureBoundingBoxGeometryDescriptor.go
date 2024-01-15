//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Metal.MTLAccelerationStructureBoundingBoxGeometryDescriptor : Metal.MTLAccelerationStructureGeometryDescriptor
*/

type MTLAccelerationStructureBoundingBoxGeometryDescriptor struct {
  *MTLAccelerationStructureGeometryDescriptor

}

func (r *MTLAccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxBufferOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxBufferOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxStride() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxStride() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureBoundingBoxGeometryDescriptor) Descriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

