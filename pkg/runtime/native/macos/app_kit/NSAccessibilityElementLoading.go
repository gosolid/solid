//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSAccessibilityElementLoading
*/

type NSAccessibilityElementLoading struct {

}

func (r *NSAccessibilityElementLoading) AccessibilityElementWithToken() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityElementLoading) AccessibilityRangeInTargetElementWithToken() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

