//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSButtonTouchBarItem : AppKit.NSTouchBarItem
*/

type NSButtonTouchBarItem struct {
  *NSTouchBarItem

}

func (r *NSButtonTouchBarItem) Action() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButtonTouchBarItem) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonTouchBarItem) Target() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButtonTouchBarItem) IsEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSButtonTouchBarItem) SetCustomizationLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonTouchBarItem) ButtonTouchBarItemWithIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButtonTouchBarItem) SetBezelColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonTouchBarItem) SetTarget() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonTouchBarItem) SetAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonTouchBarItem) SetEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonTouchBarItem) CustomizationLabel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButtonTouchBarItem) Title() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButtonTouchBarItem) SetImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSButtonTouchBarItem) Image() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSButtonTouchBarItem) BezelColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

