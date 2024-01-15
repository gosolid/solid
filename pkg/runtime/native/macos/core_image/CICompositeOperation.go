//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CICompositeOperation
*/

type CICompositeOperation struct {

}

func (r *CICompositeOperation) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CICompositeOperation) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICompositeOperation) BackgroundImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CICompositeOperation) SetBackgroundImage() error {
  return fmt.Errorf("unimplemented")
}

