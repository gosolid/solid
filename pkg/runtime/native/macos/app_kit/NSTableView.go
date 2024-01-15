//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSTableView : AppKit.NSControl
*/

type NSTableView struct {
  *NSControl
  *NSUserInterfaceValidations
  *NSTextViewDelegate
  *NSDraggingSource
  *NSAccessibilityTable
}

func (r *NSTableView) EditedRow() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) VerticalMotionCanBeginDrag() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableView) NoteHeightOfRowsWithIndexesChanged() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) RemoveRowsAtIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetAllowsMultipleSelection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetAllowsColumnSelection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) UsesAutomaticRowHeights() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableView) DeselectRow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) HighlightSelectionInClipRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetGridColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) SelectedColumn() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetDraggingDestinationFeedbackStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) UsesStaticContents() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableView) AddTableColumn() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) DidAddRowView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) SelectRowIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) InsertRowsAtIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) RegisterNib() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) GridColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) RectOfColumn() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTableView) RowAtPoint() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetGridStyleMask() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) TableColumns() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) AutosaveName() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetHighlightedTableColumn() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetAutosaveName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) ColumnWithIdentifier() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SizeToFit() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) SelectAll() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) RectOfRow() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTableView) EnumerateAvailableRowViewsUsingBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) BackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetUsesStaticContents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) SizeLastColumnToFit() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) NumberOfColumns() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetRowActionsVisible() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) RowForView() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetAllowsColumnReordering() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) EditedColumn() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) DoubleAction() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SortDescriptors() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) FloatsGroupRows() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetIndicatorImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetRowSizeStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetSortDescriptors() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) NumberOfSelectedRows() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) DataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetUsesAlternatingRowBackgroundColors() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetBackgroundColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) DeselectColumn() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) ColumnIndexesInRect() (*foundation.NSIndexSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) FrameOfCellAtColumn() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTableView) CornerView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) UsesAlternatingRowBackgroundColors() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) MoveColumn() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) CanDragRowsWithIndexes() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableView) ViewAtColumn() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) AllowsColumnReordering() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetColumnAutoresizingStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) UserInterfaceLayoutDirection() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) DragImageForRowsWithIndexes() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SelectColumnIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) RowHeight() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) RowActionsVisible() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetAllowsEmptySelection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) SelectedColumnIndexes() (*foundation.NSIndexSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetSelectionHighlightStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) DraggingDestinationFeedbackStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) InitWithFrame() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) NoteNumberOfRowsChanged() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetUserInterfaceLayoutDirection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) ReloadData() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) IndicatorImageInTableColumn() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) DrawBackgroundInClipRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) DidRemoveRowView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) RowSizeStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) AllowsMultipleSelection() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableView) MakeViewWithIdentifier() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SelectedRow() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetAllowsTypeSelect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetDropRow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) DrawRow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) UnhideRowsAtIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetHeaderView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetCornerView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) AllowsColumnResizing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableView) TableColumnWithIdentifier() (*NSTableColumn, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) Tile() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) EditColumn() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) DrawGridInClipRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetDataSource() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) ColumnAutoresizingStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) IntercellSpacing() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetIntercellSpacing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) ClickedColumn() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) AllowsTypeSelect() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableView) IsRowSelected() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetAllowsColumnResizing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) EffectiveRowSizeStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) HighlightedTableColumn() (*NSTableColumn, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) IsColumnSelected() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableView) ScrollRowToVisible() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) ReloadDataForRowIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) Style() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetAutosaveTableColumns() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) RemoveTableColumn() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) DeselectAll() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) NumberOfRows() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) HiddenRowIndexes() (*foundation.NSIndexSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetUsesAutomaticRowHeights() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetDraggingSourceOperationMask() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) BeginUpdates() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) MoveRowAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetVerticalMotionCanBeginDrag() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) RegisteredNibsByIdentifier() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) ScrollColumnToVisible() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) ColumnForView() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) AllowsEmptySelection() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetFloatsGroupRows() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) RowViewAtRow() (*NSTableRowView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) EndUpdates() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetRowHeight() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) ClickedRow() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) AllowsColumnSelection() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SelectedRowIndexes() (*foundation.NSIndexSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) AutosaveTableColumns() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableView) RowsInRect() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSTableView) ColumnAtPoint() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) HideRowsAtIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) HeaderView() (*NSTableHeaderView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableView) SetDoubleAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableView) SelectionHighlightStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) GridStyleMask() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) NumberOfSelectedColumns() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableView) EffectiveStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

