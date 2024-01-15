//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIMinimumComponent
*/

type CIMinimumComponent struct {

}

func (r *CIMinimumComponent) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIMinimumComponent) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

