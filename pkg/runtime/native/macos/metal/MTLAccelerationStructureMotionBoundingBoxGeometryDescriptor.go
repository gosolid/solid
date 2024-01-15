//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface Metal.MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor : Metal.MTLAccelerationStructureGeometryDescriptor
*/

type MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor struct {
  *MTLAccelerationStructureGeometryDescriptor

}

func (r *MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor) Descriptor() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxBuffers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxBuffers() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxStride() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxStride() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxCount() error {
  return fmt.Errorf("unimplemented")
}

