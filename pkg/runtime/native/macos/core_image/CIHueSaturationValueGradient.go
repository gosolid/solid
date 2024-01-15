//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
)

/*
protocol CoreImage.CIHueSaturationValueGradient
*/

type CIHueSaturationValueGradient struct {

}

func (r *CIHueSaturationValueGradient) Softness() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIHueSaturationValueGradient) SetDither() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIHueSaturationValueGradient) Value() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIHueSaturationValueGradient) SetValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIHueSaturationValueGradient) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIHueSaturationValueGradient) ColorSpace() (*core_graphics.CGColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIHueSaturationValueGradient) SetColorSpace() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIHueSaturationValueGradient) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIHueSaturationValueGradient) SetSoftness() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIHueSaturationValueGradient) Dither() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

