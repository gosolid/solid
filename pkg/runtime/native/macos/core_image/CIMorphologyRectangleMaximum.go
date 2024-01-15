//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIMorphologyRectangleMaximum
*/

type CIMorphologyRectangleMaximum struct {

}

func (r *CIMorphologyRectangleMaximum) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIMorphologyRectangleMaximum) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIMorphologyRectangleMaximum) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIMorphologyRectangleMaximum) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIMorphologyRectangleMaximum) Height() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIMorphologyRectangleMaximum) SetHeight() error {
  return fmt.Errorf("unimplemented")
}

