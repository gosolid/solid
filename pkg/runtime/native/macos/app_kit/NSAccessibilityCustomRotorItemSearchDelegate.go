//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSAccessibilityCustomRotorItemSearchDelegate
*/

type NSAccessibilityCustomRotorItemSearchDelegate struct {

}

func (r *NSAccessibilityCustomRotorItemSearchDelegate) Rotor() (*NSAccessibilityCustomRotorItemResult, error) {
  return nil, fmt.Errorf("unimplemented")
}

