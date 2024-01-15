//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIEightfoldReflectedTile
*/

type CIEightfoldReflectedTile struct {

}

func (r *CIEightfoldReflectedTile) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIEightfoldReflectedTile) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIEightfoldReflectedTile) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIEightfoldReflectedTile) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIEightfoldReflectedTile) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIEightfoldReflectedTile) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIEightfoldReflectedTile) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIEightfoldReflectedTile) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

