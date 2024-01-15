//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CITriangleTile
*/

type CITriangleTile struct {

}

func (r *CITriangleTile) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CITriangleTile) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITriangleTile) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CITriangleTile) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITriangleTile) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CITriangleTile) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITriangleTile) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CITriangleTile) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

