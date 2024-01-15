//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSColorPanel : AppKit.NSPanel
*/

type NSColorPanel struct {
  *NSPanel

}

func (r *NSColorPanel) ShowsAlpha() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSColorPanel) Mode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSColorPanel) SetMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPanel) DragColor() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSColorPanel) SetAccessoryView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPanel) DetachColorList() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPanel) Color() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorPanel) SetAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPanel) SetTarget() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPanel) AttachColorList() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPanel) SharedColorPanel() (*NSColorPanel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorPanel) SharedColorPanelExists() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSColorPanel) Alpha() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSColorPanel) SetPickerMask() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPanel) SetPickerMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPanel) SetContinuous() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPanel) SetShowsAlpha() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPanel) SetColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPanel) AccessoryView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorPanel) IsContinuous() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

