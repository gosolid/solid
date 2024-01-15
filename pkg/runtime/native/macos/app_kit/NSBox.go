//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSBox : AppKit.NSView
*/

type NSBox struct {
  *NSView

}

func (r *NSBox) TitlePosition() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBox) TitleFont() (*NSFont, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBox) SetContentView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBox) BorderColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBox) ContentViewMargins() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSBox) SetFrameFromContentFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBox) SetTitlePosition() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBox) ContentView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBox) IsTransparent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBox) SetCornerRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBox) FillColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBox) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBox) SizeToFit() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBox) BoxType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBox) Title() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBox) BorderRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSBox) SetContentViewMargins() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBox) SetBorderWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBox) CornerRadius() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBox) TitleRect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSBox) SetTransparent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBox) BorderWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBox) SetFillColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBox) SetTitleFont() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBox) TitleCell() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBox) SetBoxType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBox) SetBorderColor() error {
  return fmt.Errorf("unimplemented")
}

