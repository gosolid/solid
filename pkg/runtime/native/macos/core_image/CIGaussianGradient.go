//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIGaussianGradient
*/

type CIGaussianGradient struct {

}

func (r *CIGaussianGradient) Color1() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIGaussianGradient) SetColor1() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIGaussianGradient) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIGaussianGradient) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIGaussianGradient) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIGaussianGradient) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIGaussianGradient) Color0() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIGaussianGradient) SetColor0() error {
  return fmt.Errorf("unimplemented")
}

