//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIColorMonochrome
*/

type CIColorMonochrome struct {

}

func (r *CIColorMonochrome) Intensity() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIColorMonochrome) SetIntensity() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorMonochrome) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorMonochrome) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorMonochrome) Color() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorMonochrome) SetColor() error {
  return fmt.Errorf("unimplemented")
}

