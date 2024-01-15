//js:package native/ios/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface QuartzCore.CATransition : QuartzCore.CAAnimation
*/

type CATransition struct {
  *CAAnimation

}

func (r *CATransition) SetType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransition) Subtype() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransition) SetSubtype() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransition) StartProgress() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransition) SetStartProgress() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransition) EndProgress() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransition) SetEndProgress() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransition) Type() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

