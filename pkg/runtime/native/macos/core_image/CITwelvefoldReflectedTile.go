//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CITwelvefoldReflectedTile
*/

type CITwelvefoldReflectedTile struct {

}

func (r *CITwelvefoldReflectedTile) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITwelvefoldReflectedTile) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CITwelvefoldReflectedTile) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITwelvefoldReflectedTile) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CITwelvefoldReflectedTile) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITwelvefoldReflectedTile) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CITwelvefoldReflectedTile) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITwelvefoldReflectedTile) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

