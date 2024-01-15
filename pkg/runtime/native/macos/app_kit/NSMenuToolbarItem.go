//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSMenuToolbarItem : AppKit.NSToolbarItem
*/

type NSMenuToolbarItem struct {
  *NSToolbarItem

}

func (r *NSMenuToolbarItem) SetMenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuToolbarItem) ShowsIndicator() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMenuToolbarItem) SetShowsIndicator() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuToolbarItem) Menu() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

