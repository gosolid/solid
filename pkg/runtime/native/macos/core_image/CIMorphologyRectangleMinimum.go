//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIMorphologyRectangleMinimum
*/

type CIMorphologyRectangleMinimum struct {

}

func (r *CIMorphologyRectangleMinimum) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIMorphologyRectangleMinimum) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIMorphologyRectangleMinimum) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIMorphologyRectangleMinimum) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIMorphologyRectangleMinimum) Height() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIMorphologyRectangleMinimum) SetHeight() error {
  return fmt.Errorf("unimplemented")
}

