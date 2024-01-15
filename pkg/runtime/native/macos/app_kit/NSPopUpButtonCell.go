//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSPopUpButtonCell : AppKit.NSMenuItemCell
*/

type NSPopUpButtonCell struct {
  *NSMenuItemCell
  *NSMenuItemValidation
}

func (r *NSPopUpButtonCell) NumberOfItems() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) SelectedItem() (*NSMenuItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) ItemTitles() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) ArrowPosition() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) IndexOfItemWithTag() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) ItemAtIndex() (*NSMenuItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) ItemTitleAtIndex() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) SetAltersStateOfSelectedItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) PreferredEdge() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) UsesItemFromMenu() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) AddItemWithTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) IndexOfItem() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) IndexOfItemWithRepresentedObject() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) RemoveItemAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) IndexOfItemWithTarget() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) ItemWithTitle() (*NSMenuItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) SetPreferredEdge() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) SetPullsDown() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) SetArrowPosition() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) InsertItemWithTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) RemoveAllItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) SynchronizeTitleAndSelectedItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) PullsDown() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) AltersStateOfSelectedItem() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) ItemArray() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) IndexOfSelectedItem() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) IndexOfItemWithTitle() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) SelectItemWithTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) AttachPopUpWithFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) TitleOfSelectedItem() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) InitTextCell() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) SelectItemAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) SelectItemWithTag() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) DismissPopUp() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) LastItem() (*NSMenuItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) RemoveItemWithTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) PerformClickWithFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) Menu() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) AutoenablesItems() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) SetUsesItemFromMenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) AddItemsWithTitles() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) SelectItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) SetMenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButtonCell) SetAutoenablesItems() error {
  return fmt.Errorf("unimplemented")
}

