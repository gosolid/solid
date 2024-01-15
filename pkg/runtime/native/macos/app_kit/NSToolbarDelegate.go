//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
protocol AppKit.NSToolbarDelegate
*/

type NSToolbarDelegate struct {

}

func (r *NSToolbarDelegate) ToolbarImmovableItemIdentifiers() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbarDelegate) ToolbarWillAddItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarDelegate) ToolbarDidRemoveItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarDelegate) Toolbar() (*NSToolbarItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbarDelegate) ToolbarDefaultItemIdentifiers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbarDelegate) ToolbarAllowedItemIdentifiers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbarDelegate) ToolbarSelectableItemIdentifiers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

