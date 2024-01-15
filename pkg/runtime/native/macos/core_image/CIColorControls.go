//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIColorControls
*/

type CIColorControls struct {

}

func (r *CIColorControls) Contrast() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIColorControls) SetContrast() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorControls) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorControls) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorControls) Saturation() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIColorControls) SetSaturation() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorControls) Brightness() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIColorControls) SetBrightness() error {
  return fmt.Errorf("unimplemented")
}

