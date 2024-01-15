//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIRadialGradient
*/

type CIRadialGradient struct {

}

func (r *CIRadialGradient) Radius0() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRadialGradient) SetRadius0() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRadialGradient) SetRadius1() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRadialGradient) Color0() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRadialGradient) SetColor0() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRadialGradient) Color1() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRadialGradient) SetColor1() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRadialGradient) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIRadialGradient) Radius1() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRadialGradient) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

