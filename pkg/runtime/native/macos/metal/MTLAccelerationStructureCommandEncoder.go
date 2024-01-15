//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Metal.MTLAccelerationStructureCommandEncoder
*/

type MTLAccelerationStructureCommandEncoder struct {

}

func (r *MTLAccelerationStructureCommandEncoder) WaitForFence() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCommandEncoder) UseHeaps() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCommandEncoder) SampleCountersInBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCommandEncoder) BuildAccelerationStructure() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCommandEncoder) CopyAccelerationStructure() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCommandEncoder) WriteCompactedAccelerationStructureSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCommandEncoder) CopyAndCompactAccelerationStructure() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCommandEncoder) UpdateFence() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCommandEncoder) RefitAccelerationStructure() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCommandEncoder) UseResource() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCommandEncoder) UseResources() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructureCommandEncoder) UseHeap() error {
  return fmt.Errorf("unimplemented")
}

