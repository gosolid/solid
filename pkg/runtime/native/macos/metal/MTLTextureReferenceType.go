//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Metal.MTLTextureReferenceType : Metal.MTLType
*/

type MTLTextureReferenceType struct {
  *MTLType

}

func (r *MTLTextureReferenceType) TextureDataType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTextureReferenceType) TextureType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTextureReferenceType) Access() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLTextureReferenceType) IsDepthTexture() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

