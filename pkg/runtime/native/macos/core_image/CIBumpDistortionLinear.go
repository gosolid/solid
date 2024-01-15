//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIBumpDistortionLinear
*/

type CIBumpDistortionLinear struct {

}

func (r *CIBumpDistortionLinear) SetScale() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBumpDistortionLinear) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIBumpDistortionLinear) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIBumpDistortionLinear) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBumpDistortionLinear) Scale() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIBumpDistortionLinear) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBumpDistortionLinear) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIBumpDistortionLinear) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBumpDistortionLinear) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBumpDistortionLinear) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

