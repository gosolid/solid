//js:package native/ios/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface QuartzCore.CAPropertyAnimation : QuartzCore.CAAnimation
*/

type CAPropertyAnimation struct {
  *CAAnimation

}

func (r *CAPropertyAnimation) SetKeyPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAPropertyAnimation) IsAdditive() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAPropertyAnimation) IsCumulative() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAPropertyAnimation) SetCumulative() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAPropertyAnimation) AnimationWithKeyPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAPropertyAnimation) SetAdditive() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAPropertyAnimation) ValueFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAPropertyAnimation) SetValueFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAPropertyAnimation) KeyPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

