//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CICrystallize
*/

type CICrystallize struct {

}

func (r *CICrystallize) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CICrystallize) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICrystallize) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CICrystallize) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICrystallize) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CICrystallize) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

