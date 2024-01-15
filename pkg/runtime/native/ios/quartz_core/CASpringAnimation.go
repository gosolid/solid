//js:package native/ios/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface QuartzCore.CASpringAnimation : QuartzCore.CABasicAnimation
*/

type CASpringAnimation struct {
  *CABasicAnimation

}

func (r *CASpringAnimation) InitWithPerceptualDuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) Damping() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) PerceptualDuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) SetStiffness() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) SetAllowsOverdamping() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) Stiffness() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) SetDamping() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) SetInitialVelocity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) AllowsOverdamping() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) SettlingDuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) Mass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) SetMass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) InitialVelocity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) Bounce() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

