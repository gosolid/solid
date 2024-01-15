//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CITemperatureAndTint
*/

type CITemperatureAndTint struct {

}

func (r *CITemperatureAndTint) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CITemperatureAndTint) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITemperatureAndTint) Neutral() (*CIVector, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CITemperatureAndTint) SetNeutral() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITemperatureAndTint) TargetNeutral() (*CIVector, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CITemperatureAndTint) SetTargetNeutral() error {
  return fmt.Errorf("unimplemented")
}

