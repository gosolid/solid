//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "fmt"
)

/*
protocol CoreImage.CIPixellate
*/

type CIPixellate struct {

}

func (r *CIPixellate) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIPixellate) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPixellate) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIPixellate) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPixellate) Scale() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPixellate) SetScale() error {
  return fmt.Errorf("unimplemented")
}

