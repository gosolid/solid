//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSAccessibilityCheckBox
*/

type NSAccessibilityCheckBox struct {

}

func (r *NSAccessibilityCheckBox) AccessibilityValue() (*foundation.NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

