//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UITableView : UIKit.UIScrollView
*/

type UITableView struct {
  *UIScrollView

}

func (r *UITableView) DeleteSections() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SeparatorStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) ScrollToRowAtIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) RowHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetBackgroundView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetAllowsSelection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetAllowsMultipleSelectionDuringEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) RectForFooterInSection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetEstimatedSectionHeaderHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetDataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) BackgroundView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetInsetsContentViewsToSafeArea() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) RectForHeaderInSection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) RectForRowAtIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) IsPrefetchingEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) NumberOfSections() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) HasActiveDrop() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) ReloadSectionIndexTitles() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetSectionFooterHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetTableFooterView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) RectForSection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetSectionIndexMinimumDisplayRowCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) EstimatedSectionHeaderHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetAllowsSelectionDuringEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetSectionIndexTrackingBackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetSeparatorColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) ReloadRowsAtIndexPaths() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) DeselectRowAtIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) RegisterClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetSelfSizingInvalidation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetSeparatorEffect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) RemembersLastFocusedIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) EstimatedSectionFooterHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetFillerRowHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) IsEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) InsertSections() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetSectionHeaderHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) AllowsSelection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SectionIndexColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SeparatorColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetDropDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) IndexPathForCell() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) Style() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetRowHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetEstimatedRowHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) HasUncommittedUpdates() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) IndexPathsForSelectedRows() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetDragInteractionEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetSeparatorInset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetSeparatorInsetReference() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) AllowsFocusDuringEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetSelectionFollowsFocus() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) DragInteractionEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) NumberOfRowsInSection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SelectRowAtIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetEstimatedSectionFooterHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) IndexPathsForVisibleRows() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SectionIndexBackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) TableHeaderView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) CellForRowAtIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) RegisterNib() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SelfSizingInvalidation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetPrefetchingEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SeparatorInsetReference() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) BeginUpdates() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) ReloadData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) DequeueReusableCellWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetAllowsMultipleSelection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) IndexPathForSelectedRow() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) FooterViewForSection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) InsertRowsAtIndexPaths() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) AllowsSelectionDuringEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SelectionFollowsFocus() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SectionFooterHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SeparatorEffect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) ScrollToNearestSelectedRowAtScrollPosition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) MoveSection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) MoveRowAtIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) PrefetchDataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetCellLayoutMarginsFollowReadableWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) AllowsFocus() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SectionHeaderTopPadding() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetTableHeaderView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetPrefetchDataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) DropDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) CellLayoutMarginsFollowReadableWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) IndexPathForRowAtPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) ReloadSections() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) VisibleCells() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetRemembersLastFocusedIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetAllowsFocusDuringEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) ReconfigureRowsAtIndexPaths() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) DequeueReusableHeaderFooterViewWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) FillerRowHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) TableFooterView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) PerformBatchUpdates() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) EstimatedRowHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetSectionHeaderTopPadding() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) AllowsMultipleSelection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) AllowsMultipleSelectionDuringEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SectionIndexTrackingBackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) InitWithFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetSectionIndexColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) HeaderViewForSection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) EndUpdates() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) ContextMenuInteraction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetSectionIndexBackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) HasActiveDrag() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetSeparatorStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) InsetsContentViewsToSafeArea() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) DeleteRowsAtIndexPaths() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) DragDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SeparatorInset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SectionIndexMinimumDisplayRowCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) DataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetDragDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SetAllowsFocus() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) IndexPathsForRowsInRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableView) SectionHeaderHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

