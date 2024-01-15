//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface QuartzCore.CAKeyframeAnimation : QuartzCore.CAPropertyAnimation
*/

type CAKeyframeAnimation struct {
  *CAPropertyAnimation

}

func (r *CAKeyframeAnimation) Path() (*core_graphics.CGPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) SetCalculationMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) TensionValues() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) SetContinuityValues() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) SetBiasValues() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) RotationMode() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) SetRotationMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) Values() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) SetKeyTimes() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) TimingFunctions() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) ContinuityValues() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) SetValues() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) BiasValues() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) SetTensionValues() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) SetPath() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) KeyTimes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) SetTimingFunctions() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) CalculationMode() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

