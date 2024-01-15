//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CITorusLensDistortion
*/

type CITorusLensDistortion struct {

}

func (r *CITorusLensDistortion) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITorusLensDistortion) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CITorusLensDistortion) Refraction() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CITorusLensDistortion) SetRefraction() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITorusLensDistortion) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITorusLensDistortion) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CITorusLensDistortion) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITorusLensDistortion) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CITorusLensDistortion) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITorusLensDistortion) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

