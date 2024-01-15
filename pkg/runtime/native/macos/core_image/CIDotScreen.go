//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIDotScreen
*/

type CIDotScreen struct {

}

func (r *CIDotScreen) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDotScreen) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIDotScreen) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDotScreen) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDotScreen) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIDotScreen) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIDotScreen) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDotScreen) Sharpness() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIDotScreen) SetSharpness() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDotScreen) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

