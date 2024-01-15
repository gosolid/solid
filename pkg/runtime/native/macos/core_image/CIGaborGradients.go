//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIGaborGradients
*/

type CIGaborGradients struct {

}

func (r *CIGaborGradients) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIGaborGradients) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

