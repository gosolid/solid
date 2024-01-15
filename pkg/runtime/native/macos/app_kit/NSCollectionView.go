//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSCollectionView : AppKit.NSView
*/

type NSCollectionView struct {
  *NSView
  *NSDraggingSource
  *NSDraggingDestination
}

func (r *NSCollectionView) NumberOfItemsInSection() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) RegisterClass() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) DataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) SetPrefetchDataSource() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) SetContent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) IsFirstResponder() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) SetAllowsMultipleSelection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) MakeSupplementaryViewOfKind() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) IndexPathForItem() (*foundation.NSIndexPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) AllowsEmptySelection() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) SelectionIndexes() (*foundation.NSIndexSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) Content() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) SetBackgroundColors() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) SetSelectable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) SetAllowsEmptySelection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) SelectItemsAtIndexPaths() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) DeselectItemsAtIndexPaths() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) RegisterNib() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) MakeItemWithIdentifier() (*NSCollectionViewItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) IndexPathsForVisibleItems() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) InsertSections() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) DeleteSections() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) DeleteItemsAtIndexPaths() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) PrefetchDataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) FrameForItemAtIndex() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) DraggingImageForItemsAtIndexes() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) SelectionIndexPaths() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) VisibleSupplementaryViewsOfKind() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) SetDraggingSourceOperationMask() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) SetBackgroundViewScrollsWithContent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) SetCollectionViewLayout() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) ItemAtIndexPath() (*NSCollectionViewItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) IndexPathsForVisibleSupplementaryElementsOfKind() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) ReloadItemsAtIndexPaths() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) BackgroundView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) BackgroundViewScrollsWithContent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) BackgroundColors() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) SelectAll() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) ReloadSections() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) ToggleSectionCollapse() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) InsertItemsAtIndexPaths() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) CollectionViewLayout() (*NSCollectionViewLayout, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) SetSelectionIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) LayoutAttributesForItemAtIndexPath() (*NSCollectionViewLayoutAttributes, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) ItemAtIndex() (*NSCollectionViewItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) ScrollToItemsAtIndexPaths() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) AllowsMultipleSelection() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) SetDataSource() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) SetSelectionIndexPaths() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) ReloadData() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) SupplementaryViewForElementKind() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) MoveSection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) MoveItemAtIndexPath() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) NumberOfSections() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) LayoutAttributesForSupplementaryElementOfKind() (*NSCollectionViewLayoutAttributes, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) DeselectAll() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) VisibleItems() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) IndexPathForItemAtPoint() (*foundation.NSIndexPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) SetBackgroundView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) PerformBatchUpdates() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) DraggingImageForItemsAtIndexPaths() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionView) IsSelectable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

