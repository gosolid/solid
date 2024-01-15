//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSSearchToolbarItem : AppKit.NSToolbarItem
*/

type NSSearchToolbarItem struct {
  *NSToolbarItem

}

func (r *NSSearchToolbarItem) BeginSearchInteraction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSearchToolbarItem) SetSearchField() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSearchToolbarItem) SetView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSearchToolbarItem) SetResignsFirstResponderWithCancel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSearchToolbarItem) SetPreferredWidthForSearchField() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSearchToolbarItem) EndSearchInteraction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSearchToolbarItem) SearchField() (*NSSearchField, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSearchToolbarItem) View() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSearchToolbarItem) ResignsFirstResponderWithCancel() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSearchToolbarItem) PreferredWidthForSearchField() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

