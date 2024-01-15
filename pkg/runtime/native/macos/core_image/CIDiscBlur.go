//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIDiscBlur
*/

type CIDiscBlur struct {

}

func (r *CIDiscBlur) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDiscBlur) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIDiscBlur) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDiscBlur) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

