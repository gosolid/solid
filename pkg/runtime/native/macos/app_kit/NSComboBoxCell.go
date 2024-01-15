//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSComboBoxCell : AppKit.NSTextFieldCell
*/

type NSComboBoxCell struct {
  *NSTextFieldCell

}

func (r *NSComboBoxCell) ScrollItemAtIndexToVisible() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) RemoveItemWithObjectValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) UsesDataSource() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) DataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) ScrollItemAtIndexToTop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) AddItemsWithObjectValues() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) RemoveAllItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) HasVerticalScroller() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) SetHasVerticalScroller() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) IntercellSpacing() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) SetItemHeight() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) IsButtonBordered() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) CompletedString() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) IndexOfSelectedItem() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) Completes() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) SetButtonBordered() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) ItemObjectValueAtIndex() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) SetIntercellSpacing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) SetDataSource() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) SelectItemAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) ItemHeight() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) ReloadData() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) NumberOfVisibleItems() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) ObjectValueOfSelectedItem() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) InsertItemWithObjectValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) DeselectItemAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) AddItemWithObjectValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) RemoveItemAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) NumberOfItems() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) NoteNumberOfItemsChanged() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) SetNumberOfVisibleItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) SetUsesDataSource() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) SetCompletes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) IndexOfItemWithObjectValue() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) ObjectValues() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCell) SelectItemWithObjectValue() error {
  return fmt.Errorf("unimplemented")
}

