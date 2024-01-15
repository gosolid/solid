//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIRoundedRectangleGenerator
*/

type CIRoundedRectangleGenerator struct {

}

func (r *CIRoundedRectangleGenerator) SetExtent() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRoundedRectangleGenerator) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRoundedRectangleGenerator) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRoundedRectangleGenerator) Color() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRoundedRectangleGenerator) SetColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRoundedRectangleGenerator) Extent() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

