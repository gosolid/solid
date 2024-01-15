//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIMorphologyMinimum
*/

type CIMorphologyMinimum struct {

}

func (r *CIMorphologyMinimum) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIMorphologyMinimum) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIMorphologyMinimum) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIMorphologyMinimum) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

