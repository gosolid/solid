//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface QuartzCore.CABasicAnimation : QuartzCore.CAPropertyAnimation
*/

type CABasicAnimation struct {
  *CAPropertyAnimation

}

func (r *CABasicAnimation) SetToValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *CABasicAnimation) ByValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CABasicAnimation) SetByValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *CABasicAnimation) FromValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CABasicAnimation) SetFromValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *CABasicAnimation) ToValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

