//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIShadedMaterial
*/

type CIShadedMaterial struct {

}

func (r *CIShadedMaterial) Scale() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIShadedMaterial) SetScale() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIShadedMaterial) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIShadedMaterial) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIShadedMaterial) ShadingImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIShadedMaterial) SetShadingImage() error {
  return fmt.Errorf("unimplemented")
}

