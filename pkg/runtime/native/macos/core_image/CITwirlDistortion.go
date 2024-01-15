//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CITwirlDistortion
*/

type CITwirlDistortion struct {

}

func (r *CITwirlDistortion) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CITwirlDistortion) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITwirlDistortion) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CITwirlDistortion) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITwirlDistortion) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CITwirlDistortion) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITwirlDistortion) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CITwirlDistortion) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

