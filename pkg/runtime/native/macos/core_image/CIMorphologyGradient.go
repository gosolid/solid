//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIMorphologyGradient
*/

type CIMorphologyGradient struct {

}

func (r *CIMorphologyGradient) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIMorphologyGradient) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIMorphologyGradient) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIMorphologyGradient) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

