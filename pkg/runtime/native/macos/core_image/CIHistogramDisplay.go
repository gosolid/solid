//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIHistogramDisplay
*/

type CIHistogramDisplay struct {

}

func (r *CIHistogramDisplay) SetLowLimit() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIHistogramDisplay) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIHistogramDisplay) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIHistogramDisplay) Height() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIHistogramDisplay) SetHeight() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIHistogramDisplay) HighLimit() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIHistogramDisplay) SetHighLimit() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIHistogramDisplay) LowLimit() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

