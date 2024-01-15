//js:package native/macos/quartz-core
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

func (r *CASpringAnimation) AllowsOverdamping() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) SettlingDuration() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) Bounce() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) SetDamping() error {
  return fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) InitialVelocity() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) Damping() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) PerceptualDuration() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) Mass() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) Stiffness() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) SetStiffness() error {
  return fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) SetInitialVelocity() error {
  return fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) SetAllowsOverdamping() error {
  return fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) InitWithPerceptualDuration() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CASpringAnimation) SetMass() error {
  return fmt.Errorf("unimplemented")
}

