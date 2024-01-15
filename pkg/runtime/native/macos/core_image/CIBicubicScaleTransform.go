//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIBicubicScaleTransform
*/

type CIBicubicScaleTransform struct {

}

func (r *CIBicubicScaleTransform) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIBicubicScaleTransform) Scale() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIBicubicScaleTransform) AspectRatio() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIBicubicScaleTransform) SetAspectRatio() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBicubicScaleTransform) ParameterB() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIBicubicScaleTransform) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBicubicScaleTransform) SetScale() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBicubicScaleTransform) SetParameterB() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBicubicScaleTransform) ParameterC() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIBicubicScaleTransform) SetParameterC() error {
  return fmt.Errorf("unimplemented")
}

