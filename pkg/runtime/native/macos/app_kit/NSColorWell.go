//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSColorWell : AppKit.NSControl
*/

type NSColorWell struct {
  *NSControl

}

func (r *NSColorWell) PulldownTarget() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorWell) ColorWellWithStyle() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorWell) IsBordered() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSColorWell) Color() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorWell) Image() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorWell) SetImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorWell) PulldownAction() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorWell) TakeColorFrom() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorWell) IsActive() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSColorWell) SetBordered() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorWell) ColorWellStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSColorWell) SetPulldownTarget() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorWell) SupportsAlpha() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSColorWell) Deactivate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorWell) Activate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorWell) SetColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorWell) SetSupportsAlpha() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorWell) DrawWellInside() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorWell) SetColorWellStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorWell) SetPulldownAction() error {
  return fmt.Errorf("unimplemented")
}

