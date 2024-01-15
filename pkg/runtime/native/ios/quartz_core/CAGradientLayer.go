//js:package native/ios/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface QuartzCore.CAGradientLayer : QuartzCore.CALayer
*/

type CAGradientLayer struct {
  *CALayer

}

func (r *CAGradientLayer) SetEndPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAGradientLayer) Type() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAGradientLayer) Colors() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAGradientLayer) SetColors() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAGradientLayer) Locations() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAGradientLayer) SetLocations() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAGradientLayer) StartPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAGradientLayer) SetStartPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAGradientLayer) EndPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAGradientLayer) SetType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

