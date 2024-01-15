//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIGaussianBlur
*/

type CIGaussianBlur struct {

}

func (r *CIGaussianBlur) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIGaussianBlur) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIGaussianBlur) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIGaussianBlur) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

