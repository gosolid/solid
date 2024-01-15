//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIGammaAdjust
*/

type CIGammaAdjust struct {

}

func (r *CIGammaAdjust) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIGammaAdjust) Power() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIGammaAdjust) SetPower() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIGammaAdjust) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

