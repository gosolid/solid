//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Metal.MTLPrimitiveAccelerationStructureDescriptor : Metal.MTLAccelerationStructureDescriptor
*/

type MTLPrimitiveAccelerationStructureDescriptor struct {
  *MTLAccelerationStructureDescriptor

}

func (r *MTLPrimitiveAccelerationStructureDescriptor) GeometryDescriptors() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) MotionEndTime() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) Descriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) SetGeometryDescriptors() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) SetMotionEndTime() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) MotionKeyframeCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) SetMotionKeyframeCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) SetMotionStartBorderMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) MotionEndBorderMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) MotionStartTime() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) MotionStartBorderMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) SetMotionEndBorderMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPrimitiveAccelerationStructureDescriptor) SetMotionStartTime() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

