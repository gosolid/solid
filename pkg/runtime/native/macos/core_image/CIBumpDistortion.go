//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIBumpDistortion
*/

type CIBumpDistortion struct {

}

func (r *CIBumpDistortion) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIBumpDistortion) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBumpDistortion) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIBumpDistortion) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBumpDistortion) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIBumpDistortion) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBumpDistortion) Scale() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIBumpDistortion) SetScale() error {
  return fmt.Errorf("unimplemented")
}

