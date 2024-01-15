//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIFlashTransition
*/

type CIFlashTransition struct {

}

func (r *CIFlashTransition) FadeThreshold() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIFlashTransition) SetExtent() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFlashTransition) Color() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFlashTransition) StriationStrength() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIFlashTransition) SetStriationContrast() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFlashTransition) SetColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFlashTransition) SetStriationStrength() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFlashTransition) SetFadeThreshold() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFlashTransition) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFlashTransition) Extent() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CIFlashTransition) StriationContrast() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIFlashTransition) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIFlashTransition) MaxStriationRadius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIFlashTransition) SetMaxStriationRadius() error {
  return fmt.Errorf("unimplemented")
}

