//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface QuartzCore.CAGradientLayer : QuartzCore.CALayer
*/

type CAGradientLayer struct {
  *CALayer

}

func (r *CAGradientLayer) Locations() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAGradientLayer) SetStartPoint() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAGradientLayer) SetEndPoint() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAGradientLayer) Type() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAGradientLayer) SetType() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAGradientLayer) Colors() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAGradientLayer) SetColors() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAGradientLayer) EndPoint() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CAGradientLayer) SetLocations() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAGradientLayer) StartPoint() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

