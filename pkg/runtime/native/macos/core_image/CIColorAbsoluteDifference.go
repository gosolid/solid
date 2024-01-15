//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIColorAbsoluteDifference
*/

type CIColorAbsoluteDifference struct {

}

func (r *CIColorAbsoluteDifference) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorAbsoluteDifference) InputImage2() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorAbsoluteDifference) SetInputImage2() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorAbsoluteDifference) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

