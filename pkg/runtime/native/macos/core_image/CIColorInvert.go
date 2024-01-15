//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIColorInvert
*/

type CIColorInvert struct {

}

func (r *CIColorInvert) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorInvert) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

