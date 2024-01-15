//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface Metal.MTLInstanceAccelerationStructureDescriptor : Metal.MTLAccelerationStructureDescriptor
*/

type MTLInstanceAccelerationStructureDescriptor struct {
  *MTLAccelerationStructureDescriptor

}

func (r *MTLInstanceAccelerationStructureDescriptor) InstanceCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) SetInstanceCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) InstanceDescriptorType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) Descriptor() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) SetInstanceDescriptorBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) SetInstanceDescriptorStride() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) SetInstanceDescriptorType() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) MotionTransformBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) SetMotionTransformBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) MotionTransformBufferOffset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) SetMotionTransformBufferOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) InstanceDescriptorBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) SetInstanceDescriptorBufferOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) InstanceDescriptorStride() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) SetMotionTransformCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) InstanceDescriptorBufferOffset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) SetInstancedAccelerationStructures() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) MotionTransformCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLInstanceAccelerationStructureDescriptor) InstancedAccelerationStructures() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

