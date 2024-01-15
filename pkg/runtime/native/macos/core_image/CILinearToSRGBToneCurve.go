//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CILinearToSRGBToneCurve
*/

type CILinearToSRGBToneCurve struct {

}

func (r *CILinearToSRGBToneCurve) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CILinearToSRGBToneCurve) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

