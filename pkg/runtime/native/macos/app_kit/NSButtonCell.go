//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSButtonCell : AppKit.NSActionCell
*/

type NSButtonCell struct {
  *NSActionCell

}

func (r *NSButtonCell) SetKeyEquivalent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) ImageDimsWhenDisabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) GetPeriodicDelay() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) SetBezelStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) HighlightsBy() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) AlternateImage() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) ImageScaling() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) SetAlternateImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) SetImageScaling() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) IsOpaque() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) SetButtonType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) SetPeriodicDelay() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) MouseEntered() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) DrawTitle() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) SetHighlightsBy() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) SetKeyEquivalentModifierMask() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) IsTransparent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) ShowsBorderOnlyWhileMouseInside() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) SetSound() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) MouseExited() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) SetAttributedTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) AlternateTitle() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) ImagePosition() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) KeyEquivalentModifierMask() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) BackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) PerformClick() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) ShowsStateBy() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) SetTransparent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) SetImageDimsWhenDisabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) SetShowsBorderOnlyWhileMouseInside() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) BezelStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) SetShowsStateBy() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) AttributedAlternateTitle() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) SetAttributedAlternateTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) KeyEquivalent() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) InitTextCell() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) DrawBezelWithFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) Title() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) AttributedTitle() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) Sound() (*NSSound, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) SetBackgroundColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) InitImageCell() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) DrawImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) SetAlternateTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonCell) SetImagePosition() error {
  return fmt.Errorf("unimplemented")
}

