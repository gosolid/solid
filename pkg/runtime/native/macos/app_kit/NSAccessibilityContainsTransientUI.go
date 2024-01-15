//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSAccessibilityContainsTransientUI
*/

type NSAccessibilityContainsTransientUI struct {

}

func (r *NSAccessibilityContainsTransientUI) AccessibilityPerformShowAlternateUI() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityContainsTransientUI) AccessibilityPerformShowDefaultUI() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityContainsTransientUI) IsAccessibilityAlternateUIVisible() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

