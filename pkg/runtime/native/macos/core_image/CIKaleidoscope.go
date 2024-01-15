//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIKaleidoscope
*/

type CIKaleidoscope struct {

}

func (r *CIKaleidoscope) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIKaleidoscope) Count() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIKaleidoscope) SetCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIKaleidoscope) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIKaleidoscope) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIKaleidoscope) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIKaleidoscope) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIKaleidoscope) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

