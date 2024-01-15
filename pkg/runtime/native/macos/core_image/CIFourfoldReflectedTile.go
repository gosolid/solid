//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIFourfoldReflectedTile
*/

type CIFourfoldReflectedTile struct {

}

func (r *CIFourfoldReflectedTile) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIFourfoldReflectedTile) AcuteAngle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIFourfoldReflectedTile) SetAcuteAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFourfoldReflectedTile) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFourfoldReflectedTile) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFourfoldReflectedTile) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIFourfoldReflectedTile) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFourfoldReflectedTile) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFourfoldReflectedTile) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIFourfoldReflectedTile) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

