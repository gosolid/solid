//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIGlideReflectedTile
*/

type CIGlideReflectedTile struct {

}

func (r *CIGlideReflectedTile) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIGlideReflectedTile) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIGlideReflectedTile) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIGlideReflectedTile) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIGlideReflectedTile) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIGlideReflectedTile) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIGlideReflectedTile) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIGlideReflectedTile) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

