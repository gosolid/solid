//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIMorphologyMaximum
*/

type CIMorphologyMaximum struct {

}

func (r *CIMorphologyMaximum) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIMorphologyMaximum) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIMorphologyMaximum) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIMorphologyMaximum) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

