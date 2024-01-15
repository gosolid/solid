//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIDroste
*/

type CIDroste struct {

}

func (r *CIDroste) InsetPoint0() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIDroste) SetPeriodicity() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDroste) Zoom() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIDroste) SetZoom() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDroste) SetInsetPoint0() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDroste) SetRotation() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDroste) Rotation() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIDroste) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDroste) InsetPoint1() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIDroste) SetInsetPoint1() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDroste) Strands() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIDroste) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIDroste) SetStrands() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDroste) Periodicity() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

