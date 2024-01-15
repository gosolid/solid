//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIVortexDistortion
*/

type CIVortexDistortion struct {

}

func (r *CIVortexDistortion) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIVortexDistortion) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIVortexDistortion) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIVortexDistortion) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIVortexDistortion) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIVortexDistortion) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIVortexDistortion) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVortexDistortion) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

