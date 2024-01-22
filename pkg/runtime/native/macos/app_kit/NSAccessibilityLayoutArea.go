//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSAccessibilityLayoutArea
*/

type NSAccessibilityLayoutArea struct {

}

func (r *NSAccessibilityLayoutArea) AccessibilityLabel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityLayoutArea) AccessibilityChildren() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityLayoutArea) AccessibilitySelectedChildren() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityLayoutArea) AccessibilityFocusedUIElement() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}
