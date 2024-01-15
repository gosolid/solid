//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSTableHeaderView : AppKit.NSView
*/

type NSTableHeaderView struct {
  *NSView
  *NSViewToolTipOwner
}

func (r *NSTableHeaderView) DraggedDistance() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableHeaderView) ResizedColumn() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableHeaderView) HeaderRectOfColumn() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTableHeaderView) ColumnAtPoint() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableHeaderView) TableView() (*NSTableView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableHeaderView) SetTableView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableHeaderView) DraggedColumn() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

