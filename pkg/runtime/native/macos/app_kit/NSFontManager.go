//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSFontManager : objc.NSObject
*/

type NSFontManager struct {
  *objc.NSObject
  *NSMenuItemValidation
}

func (r *NSFontManager) SetFontPanelFactory() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFontManager) SetSelectedFont() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFontManager) SetFontMenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFontManager) TraitsOfFont() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) ConvertFont() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) AvailableFontFamilies() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) ConvertWeight() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) SetSelectedAttributes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFontManager) RemoveCollection() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) RemoveFontDescriptor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFontManager) SelectedFont() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) AvailableFonts() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) Action() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) FontMenu() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) WeightOfFont() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFontManager) CollectionNames() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) Target() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) FontPanel() (*NSFontPanel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) SendAction() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) ConvertAttributes() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) FontDescriptorsInCollection() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) SetEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFontManager) SetAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFontManager) AddCollection() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) SharedFontManager() (*NSFontManager, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) CurrentFontAction() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) FontWithFamily() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) LocalizedNameForFamily() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) AvailableFontNamesMatchingFontDescriptor() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) AddFontDescriptors() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFontManager) ConvertFontTraits() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) IsEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) Delegate() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) SetFontManagerFactory() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFontManager) AvailableMembersOfFontFamily() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) IsMultiple() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFontManager) SetTarget() error {
  return fmt.Errorf("unimplemented")
}

