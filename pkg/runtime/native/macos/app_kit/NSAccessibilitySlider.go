//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSAccessibilitySlider
*/

type NSAccessibilitySlider struct {

}

func (r *NSAccessibilitySlider) AccessibilityPerformDecrement() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilitySlider) AccessibilityLabel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilitySlider) AccessibilityValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilitySlider) AccessibilityPerformIncrement() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

