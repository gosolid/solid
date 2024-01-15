//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Metal.MTLComputeCommandEncoder
*/

type MTLComputeCommandEncoder struct {

}

func (r *MTLComputeCommandEncoder) UseResources() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) UseHeaps() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) MemoryBarrierWithResources() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) SetAccelerationStructure() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) SetThreadgroupMemoryLength() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) DispatchThreadgroups() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) DispatchThreads() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) UseResource() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) DispatchType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) SetBuffers() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) SetSamplerState() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) SetTextures() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) UseHeap() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) SetSamplerStates() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) ExecuteCommandsInBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) MemoryBarrierWithScope() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) SetComputePipelineState() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) SetTexture() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) UpdateFence() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) WaitForFence() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) SampleCountersInBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) SetBytes() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) SetStageInRegion() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) SetIntersectionFunctionTable() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) SetIntersectionFunctionTables() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) SetBufferOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) SetVisibleFunctionTables() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) DispatchThreadgroupsWithIndirectBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) SetVisibleFunctionTable() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) SetStageInRegionWithIndirectBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) SetBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputeCommandEncoder) SetImageblockWidth() error {
  return fmt.Errorf("unimplemented")
}

