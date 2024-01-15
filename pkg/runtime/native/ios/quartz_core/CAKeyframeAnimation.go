//js:package native/ios/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface QuartzCore.CAKeyframeAnimation : QuartzCore.CAPropertyAnimation
*/

type CAKeyframeAnimation struct {
  *CAPropertyAnimation

}

func (r *CAKeyframeAnimation) KeyTimes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) CalculationMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) BiasValues() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) SetRotationMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) Values() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) Path() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) SetPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) SetBiasValues() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) TimingFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) TensionValues() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) SetContinuityValues() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) SetTensionValues() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) RotationMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) SetValues() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) SetTimingFunctions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) SetCalculationMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) SetKeyTimes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAKeyframeAnimation) ContinuityValues() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

