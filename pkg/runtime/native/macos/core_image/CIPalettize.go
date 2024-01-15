//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIPalettize
*/

type CIPalettize struct {

}

func (r *CIPalettize) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIPalettize) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPalettize) PaletteImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIPalettize) SetPaletteImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPalettize) Perceptual() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIPalettize) SetPerceptual() error {
  return fmt.Errorf("unimplemented")
}

