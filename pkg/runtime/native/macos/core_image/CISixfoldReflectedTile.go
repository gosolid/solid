//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CISixfoldReflectedTile
*/

type CISixfoldReflectedTile struct {

}

func (r *CISixfoldReflectedTile) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CISixfoldReflectedTile) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISixfoldReflectedTile) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CISixfoldReflectedTile) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISixfoldReflectedTile) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CISixfoldReflectedTile) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISixfoldReflectedTile) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CISixfoldReflectedTile) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

