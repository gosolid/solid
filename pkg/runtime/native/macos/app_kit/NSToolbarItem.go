//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSToolbarItem : objc.NSObject
*/

type NSToolbarItem struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSToolbarItem) IsEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) IsNavigational() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) MaxSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) InitWithItemIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) ItemIdentifier() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) Target() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) Autovalidates() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) Image() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) SetNavigational() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) Toolbar() (*NSToolbar, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) MenuFormRepresentation() (*NSMenuItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) ToolTip() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) Tag() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) IsVisible() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) SetMinSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) Validate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) SetPaletteLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) SetBordered() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) MinSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) SetMaxSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) PossibleLabels() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) SetTag() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) SetTarget() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) Title() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) IsBordered() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) VisibilityPriority() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) AllowsDuplicatesInToolbar() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) PaletteLabel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) SetToolTip() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) SetEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) View() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) SetView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) SetVisibilityPriority() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) SetPossibleLabels() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) SetMenuFormRepresentation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) Action() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) SetAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) SetImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbarItem) SetAutovalidates() error {
  return fmt.Errorf("unimplemented")
}

