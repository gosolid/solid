//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSOpenSavePanelDelegate
*/

type NSOpenSavePanelDelegate struct {

}

func (r *NSOpenSavePanelDelegate) Panel() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOpenSavePanelDelegate) PanelSelectionDidChange() error {
  return fmt.Errorf("unimplemented")
}

