//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIFourfoldRotatedTile
*/

type CIFourfoldRotatedTile struct {

}

func (r *CIFourfoldRotatedTile) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIFourfoldRotatedTile) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFourfoldRotatedTile) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIFourfoldRotatedTile) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFourfoldRotatedTile) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIFourfoldRotatedTile) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFourfoldRotatedTile) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFourfoldRotatedTile) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

