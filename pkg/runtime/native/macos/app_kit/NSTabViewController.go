//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSTabViewController : AppKit.NSViewController
*/

type NSTabViewController struct {
  *NSViewController
  *NSTabViewDelegate
  *NSToolbarDelegate
}

func (r *NSTabViewController) RemoveTabViewItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabViewController) TabView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabViewController) SetTabView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabViewController) SetCanPropagateSelectedChildViewControllerTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabViewController) SelectedTabViewItemIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTabViewController) ViewDidLoad() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabViewController) Toolbar() (*NSToolbarItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTabViewController) ToolbarAllowedItemIdentifiers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTabViewController) SetTabStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabViewController) SetSelectedTabViewItemIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabViewController) InsertTabViewItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabViewController) TabStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTabViewController) TransitionOptions() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTabViewController) AddTabViewItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabViewController) TabViewItemForViewController() (*NSTabViewItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTabViewController) ToolbarDefaultItemIdentifiers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTabViewController) ToolbarSelectableItemIdentifiers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTabViewController) SetTransitionOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabViewController) CanPropagateSelectedChildViewControllerTitle() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTabViewController) TabViewItems() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTabViewController) SetTabViewItems() error {
  return fmt.Errorf("unimplemented")
}

