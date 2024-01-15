//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSPopUpButton : AppKit.NSButton
*/

type NSPopUpButton struct {
  *NSButton

}

func (r *NSPopUpButton) PullsDown() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) ItemArray() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) NumberOfItems() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) IndexOfItemWithTitle() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) ItemAtIndex() (*NSMenuItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) SelectItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) SelectItemWithTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) Menu() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) SetMenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) AutoenablesItems() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) SetPreferredEdge() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) AddItemWithTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) RemoveItemAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) IndexOfItem() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) IndexOfSelectedItem() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) IndexOfItemWithTarget() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) ItemTitles() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) ItemWithTitle() (*NSMenuItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) ItemTitleAtIndex() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) SetAutoenablesItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) PreferredEdge() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) InitWithFrame() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) RemoveAllItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) IndexOfItemWithTag() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) AddItemsWithTitles() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) SetPullsDown() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) SelectedItem() (*NSMenuItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) SelectItemAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) SynchronizeTitleAndSelectedItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) SelectedTag() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) InsertItemWithTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) RemoveItemWithTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) IndexOfItemWithRepresentedObject() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) SelectItemWithTag() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) LastItem() (*NSMenuItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopUpButton) TitleOfSelectedItem() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

