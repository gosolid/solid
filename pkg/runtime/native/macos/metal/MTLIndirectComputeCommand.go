//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Metal.MTLIndirectComputeCommand
*/

type MTLIndirectComputeCommand struct {

}

func (r *MTLIndirectComputeCommand) ConcurrentDispatchThreads() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectComputeCommand) SetBarrier() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectComputeCommand) SetImageblockWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectComputeCommand) SetComputePipelineState() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectComputeCommand) SetKernelBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectComputeCommand) ConcurrentDispatchThreadgroups() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectComputeCommand) ClearBarrier() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectComputeCommand) Reset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectComputeCommand) SetThreadgroupMemoryLength() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectComputeCommand) SetStageInRegion() error {
  return fmt.Errorf("unimplemented")
}

