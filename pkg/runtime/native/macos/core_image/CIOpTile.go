//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIOpTile
*/

type CIOpTile struct {

}

func (r *CIOpTile) Scale() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIOpTile) SetScale() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIOpTile) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIOpTile) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIOpTile) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIOpTile) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIOpTile) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIOpTile) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIOpTile) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIOpTile) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

