//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIPaletteCentroid
*/

type CIPaletteCentroid struct {

}

func (r *CIPaletteCentroid) PaletteImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIPaletteCentroid) SetPaletteImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPaletteCentroid) Perceptual() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIPaletteCentroid) SetPerceptual() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPaletteCentroid) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIPaletteCentroid) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

