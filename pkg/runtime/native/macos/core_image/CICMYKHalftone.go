//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "fmt"
)

/*
protocol CoreImage.CICMYKHalftone
*/

type CICMYKHalftone struct {

}

func (r *CICMYKHalftone) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICMYKHalftone) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CICMYKHalftone) GrayComponentReplacement() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CICMYKHalftone) UnderColorRemoval() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CICMYKHalftone) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CICMYKHalftone) SetUnderColorRemoval() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICMYKHalftone) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICMYKHalftone) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CICMYKHalftone) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CICMYKHalftone) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICMYKHalftone) Sharpness() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CICMYKHalftone) SetSharpness() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICMYKHalftone) SetGrayComponentReplacement() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICMYKHalftone) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

