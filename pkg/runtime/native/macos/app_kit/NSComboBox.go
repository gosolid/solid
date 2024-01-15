//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "fmt"
)

/*
interface AppKit.NSComboBox : AppKit.NSTextField
*/

type NSComboBox struct {
  *NSTextField

}

func (r *NSComboBox) ScrollItemAtIndexToVisible() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBox) IndexOfItemWithObjectValue() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSComboBox) ObjectValues() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComboBox) NoteNumberOfItemsChanged() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBox) AddItemWithObjectValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBox) NumberOfVisibleItems() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSComboBox) ScrollItemAtIndexToTop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBox) SelectItemAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBox) AddItemsWithObjectValues() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBox) RemoveItemAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBox) ItemObjectValueAtIndex() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComboBox) HasVerticalScroller() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSComboBox) IntercellSpacing() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSComboBox) SetNumberOfVisibleItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBox) SetButtonBordered() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBox) Completes() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSComboBox) SetCompletes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBox) DataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComboBox) SetDataSource() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBox) ObjectValueOfSelectedItem() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComboBox) SetIntercellSpacing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBox) UsesDataSource() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSComboBox) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBox) ReloadData() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBox) InsertItemWithObjectValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBox) RemoveItemWithObjectValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBox) SetHasVerticalScroller() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBox) ItemHeight() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSComboBox) RemoveAllItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBox) IndexOfSelectedItem() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSComboBox) NumberOfItems() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSComboBox) DeselectItemAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBox) SetItemHeight() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBox) SetUsesDataSource() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBox) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComboBox) SelectItemWithObjectValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBox) IsButtonBordered() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

