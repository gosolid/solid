//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSMenuItemCell : AppKit.NSButtonCell
*/

type NSMenuItemCell struct {
  *NSButtonCell

}

func (r *NSMenuItemCell) DrawTitleWithFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItemCell) DrawBorderAndBackgroundWithFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItemCell) NeedsDisplay() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMenuItemCell) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItemCell) KeyEquivalentRectForBounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSMenuItemCell) DrawKeyEquivalentWithFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItemCell) MenuItem() (*NSMenuItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItemCell) SetMenuItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItemCell) StateImageWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMenuItemCell) KeyEquivalentWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMenuItemCell) Tag() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMenuItemCell) TitleRectForBounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSMenuItemCell) StateImageRectForBounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSMenuItemCell) DrawStateImageWithFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItemCell) DrawImageWithFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItemCell) NeedsSizing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMenuItemCell) SetNeedsSizing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItemCell) SetNeedsDisplay() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItemCell) SetTag() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItemCell) CalcSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItemCell) DrawSeparatorItemWithFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMenuItemCell) ImageWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMenuItemCell) TitleWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMenuItemCell) InitTextCell() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

