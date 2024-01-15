//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSPathCell : AppKit.NSActionCell
*/

type NSPathCell struct {
  *NSActionCell
  *NSMenuItemValidation
  *NSOpenSavePanelDelegate
}

func (r *NSPathCell) PathComponentCellAtPoint() (*NSPathComponentCell, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathCell) URL() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathCell) SetBackgroundColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathCell) BackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathCell) RectOfPathComponentCell() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSPathCell) SetURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathCell) PathStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPathCell) AllowedTypes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathCell) PathComponentCells() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathCell) SetObjectValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathCell) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathCell) SetPathComponentCells() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathCell) SetPlaceholderString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathCell) SetPlaceholderAttributedString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathCell) MouseExited() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathCell) SetAllowedTypes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathCell) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathCell) PathComponentCellClass() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathCell) ClickedPathComponentCell() (*NSPathComponentCell, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathCell) MouseEntered() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathCell) PlaceholderString() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathCell) PlaceholderAttributedString() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathCell) SetPathStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathCell) DoubleAction() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathCell) SetDoubleAction() error {
  return fmt.Errorf("unimplemented")
}

