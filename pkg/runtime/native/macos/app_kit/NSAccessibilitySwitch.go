//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSAccessibilitySwitch
*/

type NSAccessibilitySwitch struct {

}

func (r *NSAccessibilitySwitch) AccessibilityValue() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilitySwitch) AccessibilityPerformIncrement() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilitySwitch) AccessibilityPerformDecrement() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

