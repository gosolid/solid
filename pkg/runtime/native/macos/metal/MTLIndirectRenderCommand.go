//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Metal.MTLIndirectRenderCommand
*/

type MTLIndirectRenderCommand struct {

}

func (r *MTLIndirectRenderCommand) SetFragmentBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectRenderCommand) DrawPatches() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectRenderCommand) DrawIndexedPatches() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectRenderCommand) DrawPrimitives() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectRenderCommand) DrawIndexedPrimitives() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectRenderCommand) Reset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectRenderCommand) SetRenderPipelineState() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectRenderCommand) SetVertexBuffer() error {
  return fmt.Errorf("unimplemented")
}

