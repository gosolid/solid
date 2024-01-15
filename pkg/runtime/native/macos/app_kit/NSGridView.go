//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSGridView : AppKit.NSView
*/

type NSGridView struct {
  *NSView

}

func (r *NSGridView) MergeCellsInHorizontalRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridView) SetRowSpacing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridView) ColumnSpacing() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGridView) RowAtIndex() (*NSGridRow, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGridView) NumberOfRows() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGridView) SetRowAlignment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridView) RowSpacing() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGridView) SetColumnSpacing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridView) CellAtColumnIndex() (*NSGridCell, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGridView) GridViewWithViews() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGridView) IndexOfRow() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGridView) CellForView() (*NSGridCell, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGridView) MoveColumnAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridView) NumberOfColumns() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGridView) SetYPlacement() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridView) GridViewWithNumberOfColumns() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGridView) AddRowWithViews() (*NSGridRow, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGridView) MoveRowAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridView) SetXPlacement() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridView) ColumnAtIndex() (*NSGridColumn, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGridView) YPlacement() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGridView) InsertColumnAtIndex() (*NSGridColumn, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGridView) RowAlignment() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGridView) InitWithFrame() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGridView) RemoveRowAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridView) AddColumnWithViews() (*NSGridColumn, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGridView) XPlacement() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGridView) IndexOfColumn() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGridView) InsertRowAtIndex() (*NSGridRow, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGridView) RemoveColumnAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGridView) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

