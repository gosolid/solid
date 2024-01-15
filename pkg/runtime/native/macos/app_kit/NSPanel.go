//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSPanel : AppKit.NSWindow
*/

type NSPanel struct {
  *NSWindow

}

func (r *NSPanel) IsFloatingPanel() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPanel) SetFloatingPanel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPanel) BecomesKeyOnlyIfNeeded() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPanel) SetBecomesKeyOnlyIfNeeded() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPanel) WorksWhenModal() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPanel) SetWorksWhenModal() error {
  return fmt.Errorf("unimplemented")
}

