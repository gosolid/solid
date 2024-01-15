//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
)

/*
protocol CoreImage.CIColorCurves
*/

type CIColorCurves struct {

}

func (r *CIColorCurves) CurvesDomain() (*CIVector, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorCurves) SetCurvesDomain() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorCurves) ColorSpace() (*core_graphics.CGColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorCurves) SetColorSpace() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorCurves) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorCurves) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorCurves) CurvesData() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorCurves) SetCurvesData() error {
  return fmt.Errorf("unimplemented")
}

