//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface QuartzCore.CAPropertyAnimation : QuartzCore.CAAnimation
*/

type CAPropertyAnimation struct {
  *CAAnimation

}

func (r *CAPropertyAnimation) SetCumulative() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAPropertyAnimation) SetValueFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAPropertyAnimation) AnimationWithKeyPath() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAPropertyAnimation) KeyPath() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAPropertyAnimation) SetAdditive() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAPropertyAnimation) ValueFunction() (*CAValueFunction, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAPropertyAnimation) SetKeyPath() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAPropertyAnimation) IsAdditive() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CAPropertyAnimation) IsCumulative() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

