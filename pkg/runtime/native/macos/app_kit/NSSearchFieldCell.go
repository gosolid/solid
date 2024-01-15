//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSSearchFieldCell : AppKit.NSTextFieldCell
*/

type NSSearchFieldCell struct {
  *NSTextFieldCell

}

func (r *NSSearchFieldCell) InitImageCell() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldCell) SearchTextRectForBounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldCell) SetRecentsAutosaveName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldCell) SendsSearchStringImmediately() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldCell) SearchButtonRectForBounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldCell) SearchButtonCell() (*NSButtonCell, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldCell) SetSearchButtonCell() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldCell) RecentsAutosaveName() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldCell) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldCell) SearchMenuTemplate() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldCell) SetSearchMenuTemplate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldCell) MaximumRecents() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldCell) SetMaximumRecents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldCell) SendsWholeSearchString() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldCell) SetSendsWholeSearchString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldCell) InitTextCell() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldCell) ResetSearchButtonCell() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldCell) ResetCancelButtonCell() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldCell) CancelButtonRectForBounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldCell) CancelButtonCell() (*NSButtonCell, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldCell) SetCancelButtonCell() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldCell) RecentSearches() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldCell) SetRecentSearches() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSearchFieldCell) SetSendsSearchStringImmediately() error {
  return fmt.Errorf("unimplemented")
}

