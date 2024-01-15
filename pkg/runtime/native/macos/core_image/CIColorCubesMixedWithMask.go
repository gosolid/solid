//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
)

/*
protocol CoreImage.CIColorCubesMixedWithMask
*/

type CIColorCubesMixedWithMask struct {

}

func (r *CIColorCubesMixedWithMask) SetCubeDimension() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorCubesMixedWithMask) Cube1Data() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorCubesMixedWithMask) ColorSpace() (*core_graphics.CGColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorCubesMixedWithMask) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorCubesMixedWithMask) Cube0Data() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorCubesMixedWithMask) SetCube0Data() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorCubesMixedWithMask) SetColorSpace() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorCubesMixedWithMask) Extrapolate() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorCubesMixedWithMask) MaskImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorCubesMixedWithMask) SetMaskImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorCubesMixedWithMask) CubeDimension() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIColorCubesMixedWithMask) SetCube1Data() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorCubesMixedWithMask) SetExtrapolate() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorCubesMixedWithMask) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

