//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLSamplerDescriptor : objc.NSObject
*/

type MTLSamplerDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLSamplerDescriptor) CompareFunction() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SupportArgumentBuffers() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) LodMinClamp() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetRAddressMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) NormalizedCoordinates() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) RAddressMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetMinFilter() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetLodMinClamp() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetSupportArgumentBuffers() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) MinFilter() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetMagFilter() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetBorderColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) LodMaxClamp() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) LodAverage() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) MagFilter() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetMipFilter() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) MaxAnisotropy() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetNormalizedCoordinates() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetLodMaxClamp() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetCompareFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) MipFilter() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SAddressMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetSAddressMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) BorderColor() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetMaxAnisotropy() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetTAddressMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) SetLodAverage() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLSamplerDescriptor) TAddressMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

