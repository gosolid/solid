//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIMix
*/

type CIMix struct {

}

func (r *CIMix) BackgroundImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIMix) SetBackgroundImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIMix) Amount() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIMix) SetAmount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIMix) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIMix) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

