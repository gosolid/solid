//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSAccessibilityRow
*/

type NSAccessibilityRow struct {

}

func (r *NSAccessibilityRow) AccessibilityIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityRow) AccessibilityDisclosureLevel() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

