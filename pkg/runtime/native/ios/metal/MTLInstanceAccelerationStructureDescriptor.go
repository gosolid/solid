//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Metal.MTLInstanceAccelerationStructureDescriptor : Metal.MTLAccelerationStructureDescriptor
*/

type MTLInstanceAccelerationStructureDescriptor struct {
  *MTLAccelerationStructureDescriptor

}

func (r *MTLInstanceAccelerationStructureDescriptor) SetInstanceCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) MotionTransformBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) InstanceDescriptorBufferOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) InstanceCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) InstanceDescriptorType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) SetInstancedAccelerationStructures() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) SetInstanceDescriptorType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) SetMotionTransformBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) SetMotionTransformBufferOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) Descriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) InstanceDescriptorStride() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) SetInstanceDescriptorStride() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) InstancedAccelerationStructures() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) MotionTransformBufferOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) MotionTransformCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) SetMotionTransformCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) InstanceDescriptorBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) SetInstanceDescriptorBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) SetInstanceDescriptorBufferOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

