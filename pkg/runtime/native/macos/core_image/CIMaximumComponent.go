//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIMaximumComponent
*/

type CIMaximumComponent struct {

}

func (r *CIMaximumComponent) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIMaximumComponent) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}
