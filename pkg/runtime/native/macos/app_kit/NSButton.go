//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSButton : AppKit.NSControl
*/

type NSButton struct {
  *NSControl
  *NSUserInterfaceValidations
  *NSAccessibilityButton
  *NSUserInterfaceCompression
}

func (r *NSButton) PerformKeyEquivalent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSButton) IsSpringLoaded() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSButton) ActiveCompressionOptions() (*NSUserInterfaceCompressionOptions, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButton) AttributedTitle() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButton) SetSpringLoaded() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) SetMaxAcceleratorLevel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) IsBordered() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSButton) IsTransparent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSButton) ImageScaling() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSButton) AttributedAlternateTitle() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButton) Sound() (*NSSound, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButton) SetShowsBorderOnlyWhileMouseInside() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) RadioButtonWithTitle() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButton) GetPeriodicDelay() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) BezelColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButton) SetImageScaling() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) KeyEquivalentModifierMask() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSButton) SetNextState() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) SetHasDestructiveAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) SetAlternateImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) ImageHugsTitle() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSButton) AllowsMixedState() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSButton) SetKeyEquivalentModifierMask() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) MaxAcceleratorLevel() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSButton) CheckboxWithTitle() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButton) ButtonWithTitle() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButton) SetPeriodicDelay() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) SetBezelStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) State() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSButton) SetButtonType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) SetAlternateTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) SetAttributedAlternateTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) SetKeyEquivalent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) ImagePosition() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSButton) SetAllowsMixedState() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) ContentTintColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButton) SetImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) AlternateImage() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButton) SetImageHugsTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) SetSymbolConfiguration() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) SetState() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) Highlight() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) Title() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButton) AlternateTitle() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButton) BezelStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSButton) SetBordered() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) SetBezelColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) Image() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButton) CompressWithPrioritizedCompressionOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) MinimumSizeWithPrioritizedCompressionOptions() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSButton) SetAttributedTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) SetContentTintColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) KeyEquivalent() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButton) ButtonWithImage() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButton) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) HasDestructiveAction() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSButton) ShowsBorderOnlyWhileMouseInside() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSButton) SetImagePosition() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) SetSound() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) SetTransparent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButton) SymbolConfiguration() (*NSImageSymbolConfiguration, error) {
  return nil, fmt.Errorf("unimplemented")
}

