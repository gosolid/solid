//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIDisintegrateWithMaskTransition
*/

type CIDisintegrateWithMaskTransition struct {

}

func (r *CIDisintegrateWithMaskTransition) ShadowRadius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIDisintegrateWithMaskTransition) SetShadowRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDisintegrateWithMaskTransition) ShadowDensity() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIDisintegrateWithMaskTransition) SetShadowDensity() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDisintegrateWithMaskTransition) ShadowOffset() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIDisintegrateWithMaskTransition) SetShadowOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDisintegrateWithMaskTransition) MaskImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIDisintegrateWithMaskTransition) SetMaskImage() error {
  return fmt.Errorf("unimplemented")
}

