//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CILineOverlay
*/

type CILineOverlay struct {

}

func (r *CILineOverlay) EdgeIntensity() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CILineOverlay) Threshold() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CILineOverlay) Contrast() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CILineOverlay) SetContrast() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILineOverlay) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILineOverlay) NRNoiseLevel() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CILineOverlay) SetNRNoiseLevel() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILineOverlay) SetEdgeIntensity() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILineOverlay) SetThreshold() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILineOverlay) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CILineOverlay) NRSharpness() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CILineOverlay) SetNRSharpness() error {
  return fmt.Errorf("unimplemented")
}

