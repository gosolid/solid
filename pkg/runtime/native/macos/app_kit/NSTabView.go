//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSTabView : AppKit.NSView
*/

type NSTabView struct {
  *NSView

}

func (r *NSTabView) NumberOfTabViewItems() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTabView) ControlTint() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTabView) TakeSelectedTabViewItemFromSender() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabView) IndexOfTabViewItem() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTabView) IndexOfTabViewItemWithIdentifier() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTabView) TabViewItemAtPoint() (*NSTabViewItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTabView) TabViewItemAtIndex() (*NSTabViewItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTabView) SetTabPosition() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabView) TabViewBorderType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTabView) SetAllowsTruncatedLabels() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabView) SelectLastTabViewItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabView) SelectPreviousTabViewItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabView) AddTabViewItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabView) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabView) InsertTabViewItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabView) ControlSize() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTabView) TabViewType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTabView) SetTabViewItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabView) DrawsBackground() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTabView) SelectTabViewItemWithIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabView) SelectedTabViewItem() (*NSTabViewItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTabView) SetFont() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabView) SelectTabViewItemAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabView) RemoveTabViewItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabView) SetDrawsBackground() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabView) SetControlSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabView) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTabView) Font() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTabView) SetTabViewBorderType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabView) MinimumSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSTabView) SetControlTint() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabView) SetTabViewType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabView) TabPosition() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTabView) ContentRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTabView) TabViewItems() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTabView) AllowsTruncatedLabels() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTabView) SelectTabViewItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabView) SelectFirstTabViewItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabView) SelectNextTabViewItem() error {
  return fmt.Errorf("unimplemented")
}

