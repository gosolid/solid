//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSSearchField : AppKit.NSTextField
*/

type NSSearchField struct {
  *NSTextField

}

func (r *NSSearchField) CancelButtonBounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSSearchField) SetRecentSearches() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSearchField) RecentsAutosaveName() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSearchField) MaximumRecents() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSearchField) SetMaximumRecents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSearchField) SetSearchMenuTemplate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSearchField) SendsSearchStringImmediately() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSearchField) SetSendsSearchStringImmediately() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSearchField) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSearchField) RecentSearches() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSearchField) SetRecentsAutosaveName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSearchField) SendsWholeSearchString() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSearchField) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSearchField) SearchTextBounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSSearchField) SearchButtonBounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSSearchField) SearchMenuTemplate() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSearchField) SetSendsWholeSearchString() error {
  return fmt.Errorf("unimplemented")
}

