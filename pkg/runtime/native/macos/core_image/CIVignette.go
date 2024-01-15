//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIVignette
*/

type CIVignette struct {

}

func (r *CIVignette) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVignette) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIVignette) Intensity() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIVignette) SetIntensity() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIVignette) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIVignette) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

