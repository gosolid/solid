//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CISRGBToneCurveToLinear
*/

type CISRGBToneCurveToLinear struct {

}

func (r *CISRGBToneCurveToLinear) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CISRGBToneCurveToLinear) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

