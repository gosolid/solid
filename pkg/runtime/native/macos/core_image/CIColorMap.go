//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIColorMap
*/

type CIColorMap struct {

}

func (r *CIColorMap) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorMap) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorMap) GradientImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorMap) SetGradientImage() error {
  return fmt.Errorf("unimplemented")
}

