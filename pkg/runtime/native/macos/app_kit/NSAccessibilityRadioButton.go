//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSAccessibilityRadioButton
*/

type NSAccessibilityRadioButton struct {

}

func (r *NSAccessibilityRadioButton) AccessibilityValue() (*foundation.NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

