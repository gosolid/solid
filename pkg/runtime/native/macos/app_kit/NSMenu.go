//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface AppKit.NSMenu : objc.NSObject
*/

type NSMenu struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSCoding
  *NSUserInterfaceItemIdentification
  *NSAppearanceCustomization
  *NSAccessibilityElement
  *NSAccessibility
}

func (r *NSMenu) InsertItemWithTitle() (*NSMenuItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenu) SetItemArray() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenu) IndexOfItem() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMenu) IndexOfItemWithSubmenu() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMenu) UserInterfaceLayoutDirection() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMenu) SetSupermenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenu) SetMinimumWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenu) InitWithTitle() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenu) MenuBarVisible() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMenu) IndexOfItemWithTarget() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMenu) HighlightedItem() (*NSMenuItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenu) RemoveAllItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenu) ItemArray() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenu) PopUpContextMenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenu) AddItemWithTitle() (*NSMenuItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenu) RemoveItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenu) IndexOfItemWithTag() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMenu) ItemWithTag() (*NSMenuItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenu) IndexOfItemWithTitle() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMenu) CancelTracking() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenu) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenu) RemoveItemAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenu) ItemAtIndex() (*NSMenuItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenu) ItemChanged() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenu) InsertItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenu) NumberOfItems() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMenu) SetUserInterfaceLayoutDirection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenu) PerformKeyEquivalent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMenu) AutoenablesItems() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMenu) MenuBarHeight() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMenu) Font() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenu) AddItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenu) PerformActionForItemAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenu) SetFont() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenu) AllowsContextMenuPlugIns() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMenu) Size() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSMenu) SetAllowsContextMenuPlugIns() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenu) IndexOfItemWithRepresentedObject() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMenu) CancelTrackingWithoutAnimation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenu) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenu) MinimumWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMenu) SetShowsStateColumn() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenu) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenu) PopUpMenuPositioningItem() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMenu) SetMenuBarVisible() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenu) SetSubmenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenu) Title() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenu) ShowsStateColumn() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMenu) ItemWithTitle() (*NSMenuItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenu) Update() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenu) Supermenu() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenu) SetAutoenablesItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenu) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

