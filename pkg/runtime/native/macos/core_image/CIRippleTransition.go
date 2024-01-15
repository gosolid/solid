//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIRippleTransition
*/

type CIRippleTransition struct {

}

func (r *CIRippleTransition) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRippleTransition) ShadingImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRippleTransition) SetShadingImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRippleTransition) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRippleTransition) SetExtent() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRippleTransition) SetScale() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRippleTransition) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIRippleTransition) Extent() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CIRippleTransition) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRippleTransition) Scale() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

