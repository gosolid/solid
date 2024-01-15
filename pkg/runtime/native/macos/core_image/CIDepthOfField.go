//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIDepthOfField
*/

type CIDepthOfField struct {

}

func (r *CIDepthOfField) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDepthOfField) SetPoint1() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDepthOfField) SetSaturation() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDepthOfField) SetUnsharpMaskRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDepthOfField) SetUnsharpMaskIntensity() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDepthOfField) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDepthOfField) Point0() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIDepthOfField) Point1() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIDepthOfField) Saturation() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIDepthOfField) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIDepthOfField) SetPoint0() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDepthOfField) UnsharpMaskIntensity() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIDepthOfField) UnsharpMaskRadius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIDepthOfField) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

