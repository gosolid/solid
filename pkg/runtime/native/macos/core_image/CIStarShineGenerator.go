//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIStarShineGenerator
*/

type CIStarShineGenerator struct {

}

func (r *CIStarShineGenerator) CrossScale() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIStarShineGenerator) CrossOpacity() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIStarShineGenerator) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIStarShineGenerator) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIStarShineGenerator) CrossWidth() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIStarShineGenerator) Epsilon() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIStarShineGenerator) SetEpsilon() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIStarShineGenerator) Color() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIStarShineGenerator) SetColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIStarShineGenerator) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIStarShineGenerator) SetCrossScale() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIStarShineGenerator) CrossAngle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIStarShineGenerator) SetCrossWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIStarShineGenerator) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIStarShineGenerator) SetCrossAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIStarShineGenerator) SetCrossOpacity() error {
  return fmt.Errorf("unimplemented")
}

