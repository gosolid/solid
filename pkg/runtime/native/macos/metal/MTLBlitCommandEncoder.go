//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Metal.MTLBlitCommandEncoder
*/

type MTLBlitCommandEncoder struct {

}

func (r *MTLBlitCommandEncoder) FillBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBlitCommandEncoder) WaitForFence() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBlitCommandEncoder) ResetTextureAccessCounters() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBlitCommandEncoder) OptimizeContentsForGPUAccess() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBlitCommandEncoder) SampleCountersInBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBlitCommandEncoder) SynchronizeTexture() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBlitCommandEncoder) GetTextureAccessCounters() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBlitCommandEncoder) CopyIndirectCommandBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBlitCommandEncoder) GenerateMipmapsForTexture() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBlitCommandEncoder) OptimizeContentsForCPUAccess() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBlitCommandEncoder) OptimizeIndirectCommandBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBlitCommandEncoder) CopyFromBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBlitCommandEncoder) CopyFromTexture() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBlitCommandEncoder) UpdateFence() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBlitCommandEncoder) ResetCommandsInBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBlitCommandEncoder) ResolveCounters() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBlitCommandEncoder) SynchronizeResource() error {
  return fmt.Errorf("unimplemented")
}

