//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSStepperCell : AppKit.NSActionCell
*/

type NSStepperCell struct {
  *NSActionCell

}

func (r *NSStepperCell) MinValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStepperCell) MaxValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStepperCell) SetMaxValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStepperCell) SetIncrement() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStepperCell) ValueWraps() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSStepperCell) Autorepeat() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSStepperCell) SetMinValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStepperCell) Increment() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStepperCell) SetValueWraps() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStepperCell) SetAutorepeat() error {
  return fmt.Errorf("unimplemented")
}

