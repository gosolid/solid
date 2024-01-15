//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSGridRow : objc.NSObject
*/

type NSGridRow struct {
  *objc.NSObject
  *foundation.NSCoding
}

func (r *NSGridRow) YPlacement() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGridRow) SetYPlacement() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridRow) SetBottomPadding() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridRow) NumberOfCells() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGridRow) TopPadding() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGridRow) CellAtIndex() (*NSGridCell, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGridRow) MergeCellsInRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridRow) RowAlignment() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGridRow) SetRowAlignment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridRow) SetHeight() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridRow) SetTopPadding() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridRow) GridView() (*NSGridView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGridRow) Height() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGridRow) BottomPadding() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGridRow) IsHidden() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSGridRow) SetHidden() error {
  return fmt.Errorf("unimplemented")
}

