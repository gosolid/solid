//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSMatrix : AppKit.NSControl
*/

type NSMatrix struct {
  *NSControl
  *NSUserInterfaceValidations
  *NSViewToolTipOwner
}

func (r *NSMatrix) RemoveColumn() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) DoubleAction() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) InitWithFrame() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) CellAtRow() (*NSCell, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) TextDidChange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SelectTextAtRow() (*NSCell, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) Cells() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SetCellSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SelectAll() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) AddRowWithCells() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) ResetCursorRects() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SetCellClass() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SelectedColumn() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) DrawsCellBackground() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SetDrawsCellBackground() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SetScrollable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) ToolTipForCell() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SetAllowsEmptySelection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) IntercellSpacing() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) BackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SetBackgroundColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) NumberOfColumns() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) AutosizesCells() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) AddColumnWithCells() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SelectText() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) AllowsEmptySelection() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) IsSelectionByRect() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SendAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SelectCellAtRow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) InsertRow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SetValidateSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) CellSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SetSelectionFrom() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) AddRow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SizeToCells() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SelectedRow() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SelectCellWithTag() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SetPrototype() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) RemoveRow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) DrawCellAtRow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) ScrollCellToVisibleAtRow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) TextShouldBeginEditing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) TextDidBeginEditing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) Prototype() (*NSCell, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SetIntercellSpacing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SetCellBackgroundColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) CellFrameAtRow() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) NumberOfRows() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) MouseDown() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) PerformKeyEquivalent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) MouseDownFlags() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) CellWithTag() (*NSCell, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) CellClass() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SetDrawsBackground() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SetDoubleAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) TextShouldEndEditing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SetAutoscroll() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SelectedCells() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) DeselectAllCells() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) PutCell() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) InsertColumn() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) AcceptsFirstMouse() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SetMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SetAutosizesCells() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SortUsingSelector() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) GetRow() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) AddColumn() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) DrawsBackground() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) IsAutoscroll() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SortUsingFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) GetNumberOfRows() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) RenewRows() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) HighlightCell() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SendDoubleAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) Mode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SelectedCell() (*NSCell, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SetAutorecalculatesCellSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SetState() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) DeselectSelectedCell() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) TextDidEndEditing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SetToolTip() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) SetSelectionByRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMatrix) CellBackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) AutorecalculatesCellSize() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMatrix) MakeCellAtRow() (*NSCell, error) {
  return nil, fmt.Errorf("unimplemented")
}

