//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIBlurredRectangleGenerator
*/

type CIBlurredRectangleGenerator struct {

}

func (r *CIBlurredRectangleGenerator) SetExtent() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBlurredRectangleGenerator) Sigma() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIBlurredRectangleGenerator) SetSigma() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBlurredRectangleGenerator) Color() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIBlurredRectangleGenerator) SetColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBlurredRectangleGenerator) Extent() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

