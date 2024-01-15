//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSBrowser : AppKit.NSControl
*/

type NSBrowser struct {
  *NSControl

}

func (r *NSBrowser) SetDraggingSourceOperationMask() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) CellClass() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetColumnResizingType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) BackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetCellClass() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) ScrollColumnsLeftBy() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) ColumnContentWidthForColumnWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) IsTitled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetAllowsTypeSelect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) ItemAtIndexPath() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) ReusesColumns() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) IndexPathForColumn() (*foundation.NSIndexPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) DoClick() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) ColumnWidthForColumnContentWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetAutohidesScroller() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) TitleOfColumn() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SelectedRowIndexesInColumn() (*foundation.NSIndexSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) LoadedCellAtRow() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SelectionIndexPath() (*foundation.NSIndexPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) IsLeafItem() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) ParentForItemsInColumn() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) Path() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SelectedRowInColumn() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) IsLoaded() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SelectionIndexPaths() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) TitleHeight() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) ReloadColumn() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetPathSeparator() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) FirstVisibleColumn() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) Tile() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetMaxVisibleColumns() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) AutohidesScroller() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SelectedCellInColumn() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) TitleFrameOfColumn() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) DefaultColumnWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) EditItemAtIndexPath() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SelectedCell() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) LastVisibleColumn() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) ItemAtRow() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) ColumnsAutosaveName() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) CanDragRowsWithIndexes() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) HasHorizontalScroller() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SelectRowIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SelectAll() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) FrameOfInsideOfColumn() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) FrameOfRow() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) DraggingImageForRowsWithIndexes() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetSendsActionOnArrowKeys() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) AllowsTypeSelect() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) LoadColumnZero() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) CellPrototype() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) MinColumnWidth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetAllowsEmptySelection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetTakesTitleFromPreviousColumn() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SelectedCells() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) RemoveSavedColumnsWithAutosaveName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) ScrollColumnsRightBy() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetDefaultColumnWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetReusesColumns() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetSelectionIndexPaths() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) ColumnResizingType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) ValidateVisibleColumns() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) ScrollColumnToVisible() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) DoDoubleClick() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetAllowsBranchSelection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) NumberOfVisibleColumns() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetHasHorizontalScroller() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetSeparatesColumns() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetRowHeight() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) ScrollRowToVisible() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) GetRow() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) AllowsEmptySelection() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SelectedColumn() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) LastColumn() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetAllowsMultipleSelection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) PathSeparator() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) PathToColumn() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) DoubleAction() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) ClickedRow() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) RowHeight() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) AddColumn() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) DrawTitleOfColumn() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) NoteHeightOfRowsWithIndexesChanged() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetDoubleAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) PrefersAllColumnUserResizing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SendAction() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SelectRow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) FrameOfColumn() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetCellPrototype() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetBackgroundColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SeparatesColumns() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetTitled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) AllowsBranchSelection() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetSelectionIndexPath() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetPrefersAllColumnUserResizing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetColumnsAutosaveName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) MaxVisibleColumns() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) AllowsMultipleSelection() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) TakesTitleFromPreviousColumn() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) ReloadDataForRowIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetPath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) WidthOfColumn() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetMinColumnWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowser) ClickedColumn() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SendsActionOnArrowKeys() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBrowser) SetLastColumn() error {
  return fmt.Errorf("unimplemented")
}

