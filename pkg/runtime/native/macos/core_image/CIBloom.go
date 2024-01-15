//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIBloom
*/

type CIBloom struct {

}

func (r *CIBloom) Intensity() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIBloom) SetIntensity() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBloom) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIBloom) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBloom) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIBloom) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

