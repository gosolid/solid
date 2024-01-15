//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CILenticularHaloGenerator
*/

type CILenticularHaloGenerator struct {

}

func (r *CILenticularHaloGenerator) StriationContrast() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CILenticularHaloGenerator) SetStriationContrast() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILenticularHaloGenerator) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CILenticularHaloGenerator) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILenticularHaloGenerator) SetHaloRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILenticularHaloGenerator) HaloWidth() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CILenticularHaloGenerator) SetColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILenticularHaloGenerator) HaloOverlap() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CILenticularHaloGenerator) StriationStrength() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CILenticularHaloGenerator) SetStriationStrength() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILenticularHaloGenerator) HaloRadius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CILenticularHaloGenerator) SetTime() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILenticularHaloGenerator) Color() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CILenticularHaloGenerator) SetHaloWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILenticularHaloGenerator) SetHaloOverlap() error {
  return fmt.Errorf("unimplemented")
}

func (r *CILenticularHaloGenerator) Time() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

