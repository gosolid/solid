//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
protocol AppKit.NSAccessibilityStaticText
*/

type NSAccessibilityStaticText struct {

}

func (r *NSAccessibilityStaticText) AccessibilityValue() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityStaticText) AccessibilityAttributedStringForRange() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityStaticText) AccessibilityVisibleCharacterRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

