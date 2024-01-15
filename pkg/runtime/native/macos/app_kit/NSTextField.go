//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSTextField : AppKit.NSControl
*/

type NSTextField struct {
  *NSControl
  *NSUserInterfaceValidations
  *NSAccessibilityNavigableStaticText
  *NSTextContent
}

func (r *NSTextField) DrawsBackground() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextField) SetBordered() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextField) SetSelectable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextField) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextField) TextShouldBeginEditing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextField) TextDidEndEditing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextField) IsBordered() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextField) SetPreferredMaxLayoutWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextField) LineBreakStrategy() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextField) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextField) SetBezelStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextField) IsEditable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextField) AcceptsFirstResponder() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextField) AllowsDefaultTighteningForTruncation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextField) SetLineBreakStrategy() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextField) TextDidBeginEditing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextField) SetPlaceholderString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextField) PlaceholderAttributedString() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextField) SetPlaceholderAttributedString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextField) SetBezeled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextField) SelectText() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextField) PlaceholderString() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextField) MaximumNumberOfLines() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextField) TextShouldEndEditing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextField) SetTextColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextField) IsBezeled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextField) BezelStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextField) SetMaximumNumberOfLines() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextField) TextDidChange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextField) SetBackgroundColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextField) TextColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextField) SetEditable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextField) IsSelectable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextField) PreferredMaxLayoutWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextField) SetAllowsDefaultTighteningForTruncation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextField) BackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextField) SetDrawsBackground() error {
  return fmt.Errorf("unimplemented")
}

