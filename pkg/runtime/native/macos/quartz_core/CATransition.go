//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface QuartzCore.CATransition : QuartzCore.CAAnimation
*/

type CATransition struct {
  *CAAnimation

}

func (r *CATransition) SetType() error {
  return fmt.Errorf("unimplemented")
}

func (r *CATransition) Subtype() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransition) StartProgress() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CATransition) EndProgress() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CATransition) Type() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransition) SetSubtype() error {
  return fmt.Errorf("unimplemented")
}

func (r *CATransition) SetStartProgress() error {
  return fmt.Errorf("unimplemented")
}

func (r *CATransition) SetEndProgress() error {
  return fmt.Errorf("unimplemented")
}

func (r *CATransition) Filter() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransition) SetFilter() error {
  return fmt.Errorf("unimplemented")
}

