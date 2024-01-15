//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSPickerTouchBarItem : AppKit.NSTouchBarItem
*/

type NSPickerTouchBarItem struct {
  *NSTouchBarItem

}

func (r *NSPickerTouchBarItem) PickerTouchBarItemWithIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) CollapsedRepresentationImage() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) SetSelectionColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) SetNumberOfOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) CustomizationLabel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) ImageAtIndex() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) LabelAtIndex() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) IsEnabledAtIndex() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) SetControlRepresentation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) SelectionMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) CollapsedRepresentationLabel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) SetCollapsedRepresentationLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) SelectedIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) SetAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) SetCustomizationLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) SetCollapsedRepresentationImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) SetSelectionMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) SetTarget() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) Action() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) SetImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) NumberOfOptions() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) ControlRepresentation() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) SetSelectedIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) Target() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) SetEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) SelectionColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPickerTouchBarItem) IsEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

