//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIGloom
*/

type CIGloom struct {

}

func (r *CIGloom) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIGloom) Intensity() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIGloom) SetIntensity() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIGloom) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIGloom) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIGloom) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

