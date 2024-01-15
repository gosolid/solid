//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICollectionView : UIKit.UIScrollView
*/

type UICollectionView struct {
  *UIScrollView

}

func (r *UICollectionView) InsertItemsAtIndexPaths() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) BeginInteractiveMovementForItemAtIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SetPrefetchingEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) DeselectItemAtIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) StartInteractiveTransitionToCollectionViewLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SupplementaryViewForElementKind() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) ScrollToItemAtIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) ReloadSections() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SetDragDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) DropDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) ReloadItemsAtIndexPaths() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) VisibleCells() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) HasActiveDrop() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) IsEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) IndexPathsForVisibleSupplementaryElementsOfKind() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SetDataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SetPrefetchDataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) AllowsFocus() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) AllowsFocusDuringEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) CancelInteractiveMovement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) NumberOfSections() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) LayoutAttributesForItemAtIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) LayoutAttributesForSupplementaryElementOfKind() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) IndexPathForItemAtPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) CellForItemAtIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) DeleteSections() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) DequeueReusableSupplementaryViewOfKind() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SelectItemAtIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) InsertSections() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) IndexPathsForSelectedItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) IndexPathsForVisibleItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SetAllowsFocus() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) UpdateInteractiveMovementTargetPosition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) DataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) IsPrefetchingEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) AllowsSelection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) HasUncommittedUpdates() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SelectionFollowsFocus() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SetAllowsMultipleSelectionDuringEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) FinishInteractiveTransition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) NumberOfItemsInSection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) DeleteItemsAtIndexPaths() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) DragDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SetAllowsMultipleSelection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) RegisterNib() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) VisibleSupplementaryViewsOfKind() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SetReorderingCadence() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) RemembersLastFocusedIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) ContextMenuInteraction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) ReorderingCadence() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SetBackgroundView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SetAllowsSelection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SetAllowsSelectionDuringEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) DragInteractionEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SetDragInteractionEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) HasActiveDrag() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) InitWithFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) DequeueConfiguredReusableSupplementaryViewWithRegistration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) CancelInteractiveTransition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) MoveItemAtIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SetSelectionFollowsFocus() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SetEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) DequeueConfiguredReusableCellWithRegistration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) IndexPathForCell() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) PerformBatchUpdates() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SetDropDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SetSelfSizingInvalidation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) AllowsSelectionDuringEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) RegisterClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) CollectionViewLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) PrefetchDataSource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SetAllowsFocusDuringEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) AllowsMultipleSelection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) DequeueReusableCellWithReuseIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) ReloadData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SetCollectionViewLayout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) ReconfigureItemsAtIndexPaths() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) BackgroundView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) MoveSection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SetRemembersLastFocusedIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) EndInteractiveMovement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) SelfSizingInvalidation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionView) AllowsMultipleSelectionDuringEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

