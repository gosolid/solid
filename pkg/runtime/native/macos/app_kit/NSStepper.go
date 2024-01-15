//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSStepper : AppKit.NSControl
*/

type NSStepper struct {
  *NSControl
  *NSAccessibilityStepper
}

func (r *NSStepper) MinValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStepper) MaxValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStepper) SetIncrement() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStepper) SetValueWraps() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStepper) SetAutorepeat() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStepper) SetMinValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStepper) SetMaxValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStepper) Increment() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStepper) ValueWraps() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSStepper) Autorepeat() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

