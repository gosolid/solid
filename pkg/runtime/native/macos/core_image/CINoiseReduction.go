//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CINoiseReduction
*/

type CINoiseReduction struct {

}

func (r *CINoiseReduction) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CINoiseReduction) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CINoiseReduction) NoiseLevel() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CINoiseReduction) SetNoiseLevel() error {
  return fmt.Errorf("unimplemented")
}

func (r *CINoiseReduction) Sharpness() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CINoiseReduction) SetSharpness() error {
  return fmt.Errorf("unimplemented")
}

