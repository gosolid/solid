//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIVibrance
*/

type CIVibrance struct {

}

func (r *CIVibrance) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIVibrance) Amount() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIVibrance) SetAmount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIVibrance) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

