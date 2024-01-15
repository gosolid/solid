//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSStepperTouchBarItem : AppKit.NSTouchBarItem
*/

type NSStepperTouchBarItem struct {
  *NSTouchBarItem

}

func (r *NSStepperTouchBarItem) StepperTouchBarItemWithIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStepperTouchBarItem) SetIncrement() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStepperTouchBarItem) Target() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStepperTouchBarItem) CustomizationLabel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStepperTouchBarItem) SetMaxValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStepperTouchBarItem) MinValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStepperTouchBarItem) SetMinValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStepperTouchBarItem) SetTarget() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStepperTouchBarItem) SetCustomizationLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStepperTouchBarItem) Increment() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStepperTouchBarItem) Value() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStepperTouchBarItem) Action() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStepperTouchBarItem) MaxValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStepperTouchBarItem) SetValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStepperTouchBarItem) SetAction() error {
  return fmt.Errorf("unimplemented")
}

