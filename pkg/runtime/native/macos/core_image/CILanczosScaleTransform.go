//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CILanczosScaleTransform
*/

type CILanczosScaleTransform struct {

}

func (r *CILanczosScaleTransform) SetScale() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILanczosScaleTransform) AspectRatio() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CILanczosScaleTransform) SetAspectRatio() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILanczosScaleTransform) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CILanczosScaleTransform) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILanczosScaleTransform) Scale() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

