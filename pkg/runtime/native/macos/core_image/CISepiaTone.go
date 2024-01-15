//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CISepiaTone
*/

type CISepiaTone struct {

}

func (r *CISepiaTone) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CISepiaTone) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISepiaTone) Intensity() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CISepiaTone) SetIntensity() error {
  return fmt.Errorf("unimplemented")
}

