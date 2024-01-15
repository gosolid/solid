//js:package native/macos/metal
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

func (r *MTLAccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxBufferOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxStride() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureBoundingBoxGeometryDescriptor) Descriptor() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureBoundingBoxGeometryDescriptor) SetBoundingBoxStride() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureBoundingBoxGeometryDescriptor) BoundingBoxBufferOffset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

