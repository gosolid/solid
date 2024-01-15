//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "fmt"
)

/*
protocol CoreImage.CIGlassLozenge
*/

type CIGlassLozenge struct {

}

func (r *CIGlassLozenge) SetRefraction() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIGlassLozenge) Point0() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIGlassLozenge) SetPoint0() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIGlassLozenge) Point1() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIGlassLozenge) SetPoint1() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIGlassLozenge) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIGlassLozenge) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIGlassLozenge) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIGlassLozenge) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIGlassLozenge) Refraction() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

