//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIHatchedScreen
*/

type CIHatchedScreen struct {

}

func (r *CIHatchedScreen) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIHatchedScreen) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIHatchedScreen) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIHatchedScreen) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIHatchedScreen) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIHatchedScreen) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIHatchedScreen) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIHatchedScreen) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIHatchedScreen) Sharpness() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIHatchedScreen) SetSharpness() error {
  return fmt.Errorf("unimplemented")
}

