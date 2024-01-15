//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSAccessibilityButton
*/

type NSAccessibilityButton struct {

}

func (r *NSAccessibilityButton) AccessibilityLabel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityButton) AccessibilityPerformPress() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

