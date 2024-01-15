//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIFourCoordinateGeometryFilter
*/

type CIFourCoordinateGeometryFilter struct {

}

func (r *CIFourCoordinateGeometryFilter) TopLeft() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIFourCoordinateGeometryFilter) SetTopLeft() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFourCoordinateGeometryFilter) TopRight() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIFourCoordinateGeometryFilter) SetBottomRight() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFourCoordinateGeometryFilter) BottomLeft() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIFourCoordinateGeometryFilter) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFourCoordinateGeometryFilter) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFourCoordinateGeometryFilter) SetBottomLeft() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFourCoordinateGeometryFilter) SetTopRight() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFourCoordinateGeometryFilter) BottomRight() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

