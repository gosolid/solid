//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Metal.MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor : Metal.MTLAccelerationStructureGeometryDescriptor
*/

type MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor struct {
  *MTLAccelerationStructureGeometryDescriptor

}

func (r *MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxStride() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxStride() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor) Descriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxBuffers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxBuffers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

