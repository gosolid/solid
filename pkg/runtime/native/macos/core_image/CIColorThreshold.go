//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIColorThreshold
*/

type CIColorThreshold struct {

}

func (r *CIColorThreshold) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorThreshold) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorThreshold) Threshold() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIColorThreshold) SetThreshold() error {
  return fmt.Errorf("unimplemented")
}

