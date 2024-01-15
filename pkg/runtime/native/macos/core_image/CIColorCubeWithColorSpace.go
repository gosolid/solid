//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
)

/*
protocol CoreImage.CIColorCubeWithColorSpace
*/

type CIColorCubeWithColorSpace struct {

}

func (r *CIColorCubeWithColorSpace) SetExtrapolate() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorCubeWithColorSpace) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorCubeWithColorSpace) SetCubeDimension() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorCubeWithColorSpace) SetColorSpace() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorCubeWithColorSpace) Extrapolate() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorCubeWithColorSpace) ColorSpace() (*core_graphics.CGColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorCubeWithColorSpace) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorCubeWithColorSpace) CubeDimension() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIColorCubeWithColorSpace) CubeData() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorCubeWithColorSpace) SetCubeData() error {
  return fmt.Errorf("unimplemented")
}

