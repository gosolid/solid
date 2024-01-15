//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSToolbar : objc.NSObject
*/

type NSToolbar struct {
  *objc.NSObject

}

func (r *NSToolbar) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbar) ValidateVisibleItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbar) AllowsUserCustomization() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSToolbar) SetSizeMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbar) Items() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbar) CenteredItemIdentifier() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbar) RemoveItemAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbar) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbar) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbar) SetAllowsUserCustomization() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbar) AllowsExtensionItems() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSToolbar) SetConfigurationFromDictionary() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbar) DisplayMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSToolbar) SetDisplayMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbar) ShowsBaselineSeparator() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSToolbar) SetShowsBaselineSeparator() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbar) VisibleItems() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbar) CenteredItemIdentifiers() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbar) AutosavesConfiguration() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSToolbar) InsertItemWithItemIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbar) SetSelectedItemIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbar) SizeMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSToolbar) ConfigurationDictionary() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbar) SelectedItemIdentifier() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbar) Identifier() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbar) SetAutosavesConfiguration() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbar) SetAllowsExtensionItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbar) InitWithIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSToolbar) RunCustomizationPalette() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbar) IsVisible() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSToolbar) SetVisible() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbar) SetCenteredItemIdentifiers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbar) SetCenteredItemIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSToolbar) CustomizationPaletteIsRunning() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

