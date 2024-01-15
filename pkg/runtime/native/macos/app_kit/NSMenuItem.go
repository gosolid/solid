//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSMenuItem : objc.NSObject
*/

type NSMenuItem struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSCoding
  *NSValidatedUserInterfaceItem
  *NSUserInterfaceItemIdentification
  *NSAccessibilityElement
  *NSAccessibility
}

func (r *NSMenuItem) OnStateImage() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) IsEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetHidden() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetKeyEquivalent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) Submenu() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) MixedStateImage() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) IsAlternate() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) Tag() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) InitWithTitle() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) HasSubmenu() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) AllowsAutomaticKeyEquivalentLocalization() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetUsesUserKeyEquivalents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetAttributedTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) IsSectionHeader() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetMixedStateImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetTag() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetRepresentedObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) View() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) IsHighlighted() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetMenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) Menu() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetKeyEquivalentModifierMask() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetState() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetOnStateImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) UsesUserKeyEquivalents() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) IsSeparatorItem() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetAllowsAutomaticKeyEquivalentMirroring() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) ParentItem() (*NSMenuItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) IndentationLevel() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetTarget() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) State() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SectionHeaderWithTitle() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) AllowsKeyEquivalentWhenHidden() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) IsHidden() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SeparatorItem() (*NSMenuItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetIndentationLevel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) ToolTip() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) Badge() (*NSMenuItemBadge, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetOffStateImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) AllowsAutomaticKeyEquivalentMirroring() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetAlternate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) IsHiddenOrHasHiddenAncestor() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetBadge() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) UserKeyEquivalent() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) Target() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) RepresentedObject() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) KeyEquivalentModifierMask() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) AttributedTitle() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) Action() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetAllowsAutomaticKeyEquivalentLocalization() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) OffStateImage() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetSubmenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) Image() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) Title() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetAllowsKeyEquivalentWhenHidden() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) SetToolTip() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItem) KeyEquivalent() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

