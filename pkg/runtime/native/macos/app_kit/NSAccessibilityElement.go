//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSAccessibilityElement
*/

type NSAccessibilityElement struct {

}

func (r *NSAccessibilityElement) AccessibilityFrame() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityElement) AccessibilityParent() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityElement) IsAccessibilityFocused() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityElement) AccessibilityIdentifier() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

