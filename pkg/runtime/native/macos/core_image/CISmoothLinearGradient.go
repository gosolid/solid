//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "fmt"
)

/*
protocol CoreImage.CISmoothLinearGradient
*/

type CISmoothLinearGradient struct {

}

func (r *CISmoothLinearGradient) SetPoint0() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISmoothLinearGradient) Point1() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CISmoothLinearGradient) SetPoint1() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISmoothLinearGradient) Color0() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CISmoothLinearGradient) SetColor0() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISmoothLinearGradient) Color1() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CISmoothLinearGradient) SetColor1() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISmoothLinearGradient) Point0() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

