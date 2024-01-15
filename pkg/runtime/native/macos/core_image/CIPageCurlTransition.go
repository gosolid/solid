//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIPageCurlTransition
*/

type CIPageCurlTransition struct {

}

func (r *CIPageCurlTransition) ShadingImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIPageCurlTransition) Extent() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CIPageCurlTransition) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPageCurlTransition) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPageCurlTransition) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPageCurlTransition) BacksideImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIPageCurlTransition) SetBacksideImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPageCurlTransition) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPageCurlTransition) SetShadingImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPageCurlTransition) SetExtent() error {
  return fmt.Errorf("unimplemented")
}

