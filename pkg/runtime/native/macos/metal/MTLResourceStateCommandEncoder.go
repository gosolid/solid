//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Metal.MTLResourceStateCommandEncoder
*/

type MTLResourceStateCommandEncoder struct {

}

func (r *MTLResourceStateCommandEncoder) UpdateTextureMapping() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLResourceStateCommandEncoder) UpdateFence() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLResourceStateCommandEncoder) WaitForFence() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLResourceStateCommandEncoder) MoveTextureMappingsFromTexture() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLResourceStateCommandEncoder) UpdateTextureMappings() error {
  return fmt.Errorf("unimplemented")
}

