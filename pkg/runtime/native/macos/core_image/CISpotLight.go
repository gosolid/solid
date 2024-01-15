//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CISpotLight
*/

type CISpotLight struct {

}

func (r *CISpotLight) SetColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISpotLight) Brightness() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CISpotLight) SetBrightness() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISpotLight) Concentration() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CISpotLight) SetLightPosition() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISpotLight) LightPointsAt() (*CIVector, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CISpotLight) SetLightPointsAt() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISpotLight) SetConcentration() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISpotLight) Color() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CISpotLight) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CISpotLight) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISpotLight) LightPosition() (*CIVector, error) {
  return nil, fmt.Errorf("unimplemented")
}

