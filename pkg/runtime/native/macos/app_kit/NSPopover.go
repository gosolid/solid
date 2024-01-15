//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSPopover : AppKit.NSResponder
*/

type NSPopover struct {
  *NSResponder
  *NSAppearanceCustomization
  *NSAccessibilityElement
  *NSAccessibility
}

func (r *NSPopover) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopover) SetContentViewController() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopover) EffectiveAppearance() (*NSAppearance, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopover) SetBehavior() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopover) SetContentSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopover) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopover) Animates() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPopover) SetPositioningRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopover) ShowRelativeToRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopover) PerformClose() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopover) SetHasFullSizeContent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopover) Appearance() (*NSAppearance, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopover) Behavior() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPopover) IsShown() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPopover) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopover) Close() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopover) IsDetached() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPopover) PositioningRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSPopover) SetAppearance() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopover) SetAnimates() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPopover) ContentViewController() (*NSViewController, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopover) ContentSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSPopover) HasFullSizeContent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPopover) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPopover) ShowRelativeToToolbarItem() error {
  return fmt.Errorf("unimplemented")
}

