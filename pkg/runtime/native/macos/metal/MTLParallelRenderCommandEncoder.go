//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Metal.MTLParallelRenderCommandEncoder
*/

type MTLParallelRenderCommandEncoder struct {

}

func (r *MTLParallelRenderCommandEncoder) SetColorStoreAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLParallelRenderCommandEncoder) SetDepthStoreAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLParallelRenderCommandEncoder) SetStencilStoreAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLParallelRenderCommandEncoder) SetColorStoreActionOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLParallelRenderCommandEncoder) SetDepthStoreActionOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLParallelRenderCommandEncoder) SetStencilStoreActionOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLParallelRenderCommandEncoder) RenderCommandEncoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

