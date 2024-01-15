//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIDither
*/

type CIDither struct {

}

func (r *CIDither) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIDither) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDither) Intensity() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIDither) SetIntensity() error {
  return fmt.Errorf("unimplemented")
}

