//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSAccessibilityStepper
*/

type NSAccessibilityStepper struct {

}

func (r *NSAccessibilityStepper) AccessibilityLabel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityStepper) AccessibilityPerformIncrement() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityStepper) AccessibilityPerformDecrement() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityStepper) AccessibilityValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

