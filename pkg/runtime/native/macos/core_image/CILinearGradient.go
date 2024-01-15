//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CILinearGradient
*/

type CILinearGradient struct {

}

func (r *CILinearGradient) SetPoint1() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILinearGradient) Color0() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CILinearGradient) SetColor0() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILinearGradient) Color1() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CILinearGradient) SetColor1() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILinearGradient) Point0() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CILinearGradient) SetPoint0() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILinearGradient) Point1() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

