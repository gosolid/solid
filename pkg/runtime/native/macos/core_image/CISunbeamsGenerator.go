//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CISunbeamsGenerator
*/

type CISunbeamsGenerator struct {

}

func (r *CISunbeamsGenerator) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CISunbeamsGenerator) Color() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CISunbeamsGenerator) SetMaxStriationRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISunbeamsGenerator) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISunbeamsGenerator) SetSunRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISunbeamsGenerator) Time() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CISunbeamsGenerator) SunRadius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CISunbeamsGenerator) MaxStriationRadius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CISunbeamsGenerator) SetStriationStrength() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISunbeamsGenerator) SetColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISunbeamsGenerator) StriationStrength() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CISunbeamsGenerator) StriationContrast() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CISunbeamsGenerator) SetStriationContrast() error {
  return fmt.Errorf("unimplemented")
}

func (r *CISunbeamsGenerator) SetTime() error {
  return fmt.Errorf("unimplemented")
}

