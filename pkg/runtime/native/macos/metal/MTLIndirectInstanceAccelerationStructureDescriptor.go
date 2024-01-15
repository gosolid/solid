//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Metal.MTLIndirectInstanceAccelerationStructureDescriptor : Metal.MTLAccelerationStructureDescriptor
*/

type MTLIndirectInstanceAccelerationStructureDescriptor struct {
  *MTLAccelerationStructureDescriptor

}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) MaxInstanceCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorType() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) InstanceCountBufferOffset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) SetInstanceCountBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) MotionTransformBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) SetMotionTransformBufferOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) MotionTransformCountBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorBufferOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) SetMaxInstanceCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) InstanceCountBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) MaxMotionTransformCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) SetMotionTransformCountBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) MotionTransformCountBufferOffset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorBufferOffset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) SetInstanceCountBufferOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) SetMotionTransformBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) SetMaxMotionTransformCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) Descriptor() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) SetInstanceDescriptorStride() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) InstanceDescriptorStride() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) SetMotionTransformCountBufferOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectInstanceAccelerationStructureDescriptor) MotionTransformBufferOffset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

