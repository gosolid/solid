//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIPinchDistortion
*/

type CIPinchDistortion struct {

}

func (r *CIPinchDistortion) Scale() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPinchDistortion) SetScale() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPinchDistortion) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIPinchDistortion) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPinchDistortion) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIPinchDistortion) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPinchDistortion) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPinchDistortion) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

