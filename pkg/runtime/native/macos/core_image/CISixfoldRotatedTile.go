//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CISixfoldRotatedTile
*/

type CISixfoldRotatedTile struct {

}

func (r *CISixfoldRotatedTile) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CISixfoldRotatedTile) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISixfoldRotatedTile) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CISixfoldRotatedTile) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISixfoldRotatedTile) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CISixfoldRotatedTile) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISixfoldRotatedTile) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CISixfoldRotatedTile) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

