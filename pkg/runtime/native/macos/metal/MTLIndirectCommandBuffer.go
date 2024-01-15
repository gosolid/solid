//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Metal.MTLIndirectCommandBuffer
*/

type MTLIndirectCommandBuffer struct {

}

func (r *MTLIndirectCommandBuffer) ResetWithRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBuffer) IndirectRenderCommandAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBuffer) IndirectComputeCommandAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBuffer) Size() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLIndirectCommandBuffer) GpuResourceID() (MTLResourceID, error) {
  return MTLResourceID{}, fmt.Errorf("unimplemented")
}

