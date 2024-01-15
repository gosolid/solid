//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIExposureAdjust
*/

type CIExposureAdjust struct {

}

func (r *CIExposureAdjust) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIExposureAdjust) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIExposureAdjust) EV() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIExposureAdjust) SetEV() error {
  return fmt.Errorf("unimplemented")
}

