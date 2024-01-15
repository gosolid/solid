//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSFontPanel : AppKit.NSPanel
*/

type NSFontPanel struct {
  *NSPanel

}

func (r *NSFontPanel) SharedFontPanelExists() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFontPanel) WorksWhenModal() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFontPanel) SetWorksWhenModal() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFontPanel) SetEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFontPanel) ReloadDefaultFontFamilies() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFontPanel) PanelConvertFont() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontPanel) SharedFontPanel() (*NSFontPanel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontPanel) AccessoryView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFontPanel) SetAccessoryView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFontPanel) IsEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFontPanel) SetPanelFont() error {
  return fmt.Errorf("unimplemented")
}

