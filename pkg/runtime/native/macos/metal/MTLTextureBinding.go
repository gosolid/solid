//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Metal.MTLTextureBinding
*/

type MTLTextureBinding struct {

}

func (r *MTLTextureBinding) IsDepthTexture() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLTextureBinding) ArrayLength() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTextureBinding) TextureType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTextureBinding) TextureDataType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

