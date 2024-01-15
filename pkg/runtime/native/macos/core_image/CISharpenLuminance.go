//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CISharpenLuminance
*/

type CISharpenLuminance struct {

}

func (r *CISharpenLuminance) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISharpenLuminance) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CISharpenLuminance) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISharpenLuminance) Sharpness() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CISharpenLuminance) SetSharpness() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISharpenLuminance) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

