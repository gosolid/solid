//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIHoleDistortion
*/

type CIHoleDistortion struct {

}

func (r *CIHoleDistortion) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIHoleDistortion) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIHoleDistortion) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIHoleDistortion) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIHoleDistortion) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIHoleDistortion) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

