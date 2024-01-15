//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIPerspectiveTile
*/

type CIPerspectiveTile struct {

}

func (r *CIPerspectiveTile) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPerspectiveTile) TopLeft() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIPerspectiveTile) BottomRight() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIPerspectiveTile) SetBottomRight() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPerspectiveTile) SetBottomLeft() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPerspectiveTile) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIPerspectiveTile) SetTopLeft() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPerspectiveTile) TopRight() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIPerspectiveTile) SetTopRight() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPerspectiveTile) BottomLeft() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

