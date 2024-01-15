//js:package native/ios/metal
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

func (r *MTLTextureReferenceType) TextureDataType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureReferenceType) TextureType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureReferenceType) Access() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTextureReferenceType) IsDepthTexture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

