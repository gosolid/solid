//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSStatusBarButton : AppKit.NSButton
*/

type NSStatusBarButton struct {
  *NSButton

}

func (r *NSStatusBarButton) AppearsDisabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSStatusBarButton) SetAppearsDisabled() error {
  return fmt.Errorf("unimplemented")
}

