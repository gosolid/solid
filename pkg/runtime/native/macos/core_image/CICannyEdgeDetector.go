//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CICannyEdgeDetector
*/

type CICannyEdgeDetector struct {

}

func (r *CICannyEdgeDetector) SetGaussianSigma() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICannyEdgeDetector) SetPerceptual() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICannyEdgeDetector) ThresholdHigh() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CICannyEdgeDetector) SetThresholdHigh() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICannyEdgeDetector) ThresholdLow() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CICannyEdgeDetector) SetThresholdLow() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICannyEdgeDetector) HysteresisPasses() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CICannyEdgeDetector) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CICannyEdgeDetector) SetHysteresisPasses() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICannyEdgeDetector) GaussianSigma() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CICannyEdgeDetector) Perceptual() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CICannyEdgeDetector) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

