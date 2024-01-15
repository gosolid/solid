//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol CoreImage.CIColorCube
*/

type CIColorCube struct {

}

func (r *CIColorCube) CubeDimension() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIColorCube) SetCubeDimension() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorCube) CubeData() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorCube) SetCubeData() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorCube) Extrapolate() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorCube) SetExtrapolate() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorCube) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorCube) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

