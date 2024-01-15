//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSText : AppKit.NSView
*/

type NSText struct {
  *NSView
  *NSChangeSpelling
  *NSIgnoreMisspelledWords
}

func (r *NSText) Alignment() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSText) InitWithFrame() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSText) SetFont() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) Copy() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) SetUsesFontPanel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) SetVerticallyResizable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) IsRichText() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSText) SetHorizontallyResizable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) CheckSpelling() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) SelectedRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSText) SetAlignment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) SetBaseWritingDirection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) MinSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSText) PasteFont() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSText) WriteRTFDToFile() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSText) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSText) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) SetImportsGraphics() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) IsFieldEditor() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSText) SetDrawsBackground() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) BackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSText) CopyFont() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) SetFieldEditor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) IsVerticallyResizable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSText) RTFFromRange() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSText) AlignCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) Subscript() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) String() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSText) SetString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) UsesFontPanel() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSText) SetMaxSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) SetMinSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) RTFDFromRange() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSText) CopyRuler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) SetSelectable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) IsSelectable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSText) ReplaceCharactersInRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) ReadRTFDFromFile() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSText) SizeToFit() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) Delete() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) ChangeFont() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) AlignLeft() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) ToggleRuler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) IsRulerVisible() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSText) TextColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSText) Cut() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) Superscript() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) SetRichText() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) SetSelectedRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) SetTextColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) SetBackgroundColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) Font() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSText) BaseWritingDirection() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSText) PasteRuler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) SelectAll() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) Unscript() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) ShowGuessPanel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) DrawsBackground() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSText) MaxSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSText) IsHorizontallyResizable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSText) ScrollRangeToVisible() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) ImportsGraphics() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSText) Paste() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) AlignRight() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) Underline() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSText) IsEditable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSText) SetEditable() error {
  return fmt.Errorf("unimplemented")
}

