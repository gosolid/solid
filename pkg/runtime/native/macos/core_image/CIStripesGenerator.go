//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIStripesGenerator
*/

type CIStripesGenerator struct {

}

func (r *CIStripesGenerator) SetColor0() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIStripesGenerator) Color1() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIStripesGenerator) SetColor1() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIStripesGenerator) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIStripesGenerator) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIStripesGenerator) Sharpness() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIStripesGenerator) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIStripesGenerator) Color0() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIStripesGenerator) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIStripesGenerator) SetSharpness() error {
  return fmt.Errorf("unimplemented")
}

