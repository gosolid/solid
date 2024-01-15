//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLSamplerDescriptor : objc.NSObject
*/

type MTLSamplerDescriptor struct {
  *objc.NSObject

}

func (r *MTLSamplerDescriptor) SetMinFilter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetTAddressMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) LodMaxClamp() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetLodAverage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) MinFilter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) MagFilter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) LodAverage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) Label() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetBorderColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) NormalizedCoordinates() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetNormalizedCoordinates() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetSupportArgumentBuffers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetSAddressMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) TAddressMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) CompareFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) LodMinClamp() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetLodMinClamp() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetLabel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetMagFilter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) MipFilter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetCompareFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SupportArgumentBuffers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) BorderColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetMipFilter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) MaxAnisotropy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetMaxAnisotropy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetRAddressMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SAddressMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) RAddressMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetLodMaxClamp() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

