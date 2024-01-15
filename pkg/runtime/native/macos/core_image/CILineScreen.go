//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CILineScreen
*/

type CILineScreen struct {

}

func (r *CILineScreen) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CILineScreen) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILineScreen) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CILineScreen) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILineScreen) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CILineScreen) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILineScreen) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CILineScreen) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILineScreen) Sharpness() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CILineScreen) SetSharpness() error {
  return fmt.Errorf("unimplemented")
}

