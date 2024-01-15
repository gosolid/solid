//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "fmt"
)

/*
protocol CoreImage.CIPageCurlWithShadowTransition
*/

type CIPageCurlWithShadowTransition struct {

}

func (r *CIPageCurlWithShadowTransition) SetBacksideImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPageCurlWithShadowTransition) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPageCurlWithShadowTransition) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPageCurlWithShadowTransition) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPageCurlWithShadowTransition) ShadowAmount() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPageCurlWithShadowTransition) SetShadowExtent() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPageCurlWithShadowTransition) BacksideImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIPageCurlWithShadowTransition) SetExtent() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPageCurlWithShadowTransition) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPageCurlWithShadowTransition) ShadowSize() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPageCurlWithShadowTransition) ShadowExtent() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CIPageCurlWithShadowTransition) Extent() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CIPageCurlWithShadowTransition) SetShadowSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPageCurlWithShadowTransition) SetShadowAmount() error {
  return fmt.Errorf("unimplemented")
}

