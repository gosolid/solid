//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSSwitch : AppKit.NSControl
*/

type NSSwitch struct {
  *NSControl
  *NSAccessibilitySwitch
}

func (r *NSSwitch) State() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSwitch) SetState() error {
  return fmt.Errorf("unimplemented")
}

