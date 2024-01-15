//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSToolbarItemGroup : AppKit.NSToolbarItem
*/

type NSToolbarItemGroup struct {
  *NSToolbarItem

}

func (r *NSToolbarItemGroup) GroupWithItemIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItemGroup) IsSelectedAtIndex() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItemGroup) SetSubitems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarItemGroup) ControlRepresentation() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItemGroup) SetControlRepresentation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarItemGroup) SelectedIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItemGroup) SetSelectedIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarItemGroup) SetSelected() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarItemGroup) Subitems() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItemGroup) SelectionMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItemGroup) SetSelectionMode() error {
  return fmt.Errorf("unimplemented")
}

