//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSComboButton : AppKit.NSControl
*/

type NSComboButton struct {
  *NSControl

}

func (r *NSComboButton) ComboButtonWithImage() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComboButton) Image() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComboButton) SetImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboButton) ImageScaling() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSComboButton) Menu() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComboButton) SetMenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboButton) ComboButtonWithTitle() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComboButton) Title() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComboButton) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboButton) SetImageScaling() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboButton) Style() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSComboButton) SetStyle() error {
  return fmt.Errorf("unimplemented")
}

