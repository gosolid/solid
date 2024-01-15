//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CISobelGradients
*/

type CISobelGradients struct {

}

func (r *CISobelGradients) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CISobelGradients) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

