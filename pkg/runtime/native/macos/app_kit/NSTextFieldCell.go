//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSTextFieldCell : AppKit.NSActionCell
*/

type NSTextFieldCell struct {
  *NSActionCell

}

func (r *NSTextFieldCell) BezelStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextFieldCell) SetBezelStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFieldCell) SetAllowedInputSourceLocales() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFieldCell) PlaceholderAttributedString() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextFieldCell) SetUpFieldEditorAttributes() (*NSText, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextFieldCell) SetWantsNotificationForMarkedText() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFieldCell) BackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextFieldCell) SetBackgroundColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFieldCell) SetPlaceholderString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFieldCell) AllowedInputSourceLocales() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextFieldCell) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextFieldCell) DrawsBackground() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextFieldCell) SetDrawsBackground() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFieldCell) SetTextColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFieldCell) SetPlaceholderAttributedString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFieldCell) InitTextCell() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextFieldCell) InitImageCell() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextFieldCell) TextColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextFieldCell) PlaceholderString() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

