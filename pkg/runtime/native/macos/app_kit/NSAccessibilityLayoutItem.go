//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSAccessibilityLayoutItem
*/

type NSAccessibilityLayoutItem struct {

}

func (r *NSAccessibilityLayoutItem) SetAccessibilityFrame() error {
  return fmt.Errorf("unimplemented")
}

