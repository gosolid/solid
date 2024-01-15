//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "fmt"
)

/*
protocol CoreImage.CITriangleKaleidoscope
*/

type CITriangleKaleidoscope struct {

}

func (r *CITriangleKaleidoscope) SetPoint() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITriangleKaleidoscope) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CITriangleKaleidoscope) Point() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CITriangleKaleidoscope) SetSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITriangleKaleidoscope) Rotation() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CITriangleKaleidoscope) SetRotation() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITriangleKaleidoscope) Decay() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CITriangleKaleidoscope) SetDecay() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITriangleKaleidoscope) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITriangleKaleidoscope) Size() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

