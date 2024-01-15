//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSColorPickerTouchBarItem : AppKit.NSTouchBarItem
*/

type NSColorPickerTouchBarItem struct {
  *NSTouchBarItem

}

func (r *NSColorPickerTouchBarItem) ColorList() (*NSColorList, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorPickerTouchBarItem) SetTarget() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPickerTouchBarItem) SetAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPickerTouchBarItem) SetEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPickerTouchBarItem) Color() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorPickerTouchBarItem) Action() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorPickerTouchBarItem) IsEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSColorPickerTouchBarItem) SetShowsAlpha() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPickerTouchBarItem) ShowsAlpha() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSColorPickerTouchBarItem) AllowedColorSpaces() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorPickerTouchBarItem) SetAllowedColorSpaces() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPickerTouchBarItem) SetColorList() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPickerTouchBarItem) Target() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorPickerTouchBarItem) SetColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPickerTouchBarItem) TextColorPickerWithIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorPickerTouchBarItem) StrokeColorPickerWithIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorPickerTouchBarItem) CustomizationLabel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorPickerTouchBarItem) SetCustomizationLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPickerTouchBarItem) ColorPickerWithIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

