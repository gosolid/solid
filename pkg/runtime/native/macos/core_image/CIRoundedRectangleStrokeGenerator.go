//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIRoundedRectangleStrokeGenerator
*/

type CIRoundedRectangleStrokeGenerator struct {

}

func (r *CIRoundedRectangleStrokeGenerator) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRoundedRectangleStrokeGenerator) Color() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRoundedRectangleStrokeGenerator) SetColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRoundedRectangleStrokeGenerator) Extent() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CIRoundedRectangleStrokeGenerator) SetExtent() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRoundedRectangleStrokeGenerator) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRoundedRectangleStrokeGenerator) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRoundedRectangleStrokeGenerator) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

