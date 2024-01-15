//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol AppKit.NSAccessibilityNavigableStaticText
*/

type NSAccessibilityNavigableStaticText struct {

}

func (r *NSAccessibilityNavigableStaticText) AccessibilityStringForRange() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityNavigableStaticText) AccessibilityLineForIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityNavigableStaticText) AccessibilityRangeForLine() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityNavigableStaticText) AccessibilityFrameForRange() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

