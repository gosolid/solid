//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface Metal.MTLPrimitiveAccelerationStructureDescriptor : Metal.MTLAccelerationStructureDescriptor
*/

type MTLPrimitiveAccelerationStructureDescriptor struct {
  *MTLAccelerationStructureDescriptor

}

func (r *MTLPrimitiveAccelerationStructureDescriptor) SetGeometryDescriptors() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) MotionKeyframeCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) SetMotionEndBorderMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) MotionStartTime() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) SetMotionStartTime() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) SetMotionStartBorderMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) MotionEndTime() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) SetMotionEndTime() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) SetMotionKeyframeCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) Descriptor() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) GeometryDescriptors() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) MotionStartBorderMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) MotionEndBorderMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

