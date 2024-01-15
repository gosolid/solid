//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "fmt"
)

/*
interface QuartzCore.CAEmitterLayer : QuartzCore.CALayer
*/

type CAEmitterLayer struct {
  *CALayer

}

func (r *CAEmitterLayer) Seed() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) SetEmitterSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) RenderMode() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) SetSpin() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) EmitterZPosition() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) SetLifetime() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) EmitterPosition() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) SetEmitterPosition() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) PreservesDepth() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) Velocity() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) EmitterCells() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) BirthRate() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) SetBirthRate() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) Scale() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) EmitterSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) EmitterMode() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) SetSeed() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) EmitterDepth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) SetPreservesDepth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) Spin() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) SetScale() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) SetEmitterCells() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) Lifetime() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) SetRenderMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) EmitterShape() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) SetEmitterShape() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) SetVelocity() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) SetEmitterZPosition() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) SetEmitterDepth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAEmitterLayer) SetEmitterMode() error {
  return fmt.Errorf("unimplemented")
}

