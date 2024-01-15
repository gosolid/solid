//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "fmt"
)

/*
interface AppKit.NSOutlineView : AppKit.NSTableView
*/

type NSOutlineView struct {
  *NSTableView
  *NSAccessibilityOutline
}

func (r *NSOutlineView) SetIndentationPerLevel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) IndentationMarkerFollowsCell() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) NumberOfChildrenOfItem() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) ReloadItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) InsertItemsAtIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) ChildIndexForItem() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) ItemAtRow() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) LevelForItem() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) UserInterfaceLayoutDirection() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) RemoveRowsAtIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) OutlineTableColumn() (*NSTableColumn, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) IndentationPerLevel() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) Child() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) ExpandItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) RowForItem() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) MoveItemAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) AutoresizesOutlineColumn() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) SetAutoresizesOutlineColumn() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) AutosaveExpandedItems() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) SetAutosaveExpandedItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) ParentForItem() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) LevelForRow() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) SetDataSource() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) SetOutlineTableColumn() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) StronglyReferencesItems() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) CollapseItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) FrameOfOutlineCellAtRow() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) MoveRowAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) DataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) IsExpandable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) SetDropItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) SetIndentationMarkerFollowsCell() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) SetStronglyReferencesItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) ShouldCollapseAutoExpandedItemsForDeposited() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) IsItemExpanded() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) RemoveItemsAtIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) InsertRowsAtIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineView) SetUserInterfaceLayoutDirection() error {
  return fmt.Errorf("unimplemented")
}

