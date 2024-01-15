//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSTouchBar : objc.NSObject
*/

type NSTouchBar struct {
  *objc.NSObject
  *foundation.NSCoding
}

func (r *NSTouchBar) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTouchBar) CustomizationAllowedItemIdentifiers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTouchBar) CustomizationRequiredItemIdentifiers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTouchBar) SetDefaultItemIdentifiers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTouchBar) PrincipalItemIdentifier() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTouchBar) TemplateItems() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTouchBar) CustomizationIdentifier() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTouchBar) SetCustomizationIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTouchBar) ItemIdentifiers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTouchBar) SetTemplateItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTouchBar) SetEscapeKeyReplacementItemIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTouchBar) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTouchBar) IsVisible() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTouchBar) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTouchBar) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTouchBar) ItemForIdentifier() (*NSTouchBarItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTouchBar) SetCustomizationRequiredItemIdentifiers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTouchBar) EscapeKeyReplacementItemIdentifier() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTouchBar) IsAutomaticCustomizeTouchBarMenuItemEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTouchBar) SetCustomizationAllowedItemIdentifiers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTouchBar) DefaultItemIdentifiers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTouchBar) SetPrincipalItemIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTouchBar) SetAutomaticCustomizeTouchBarMenuItemEnabled() error {
  return fmt.Errorf("unimplemented")
}

