//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UICollectionViewLayoutInvalidationContext : objc.NSObject
*/

type UICollectionViewLayoutInvalidationContext struct {
  *objc.NSObject

}

func (r *UICollectionViewLayoutInvalidationContext) InvalidateItemsAtIndexPaths() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutInvalidationContext) InvalidateSupplementaryElementsOfKind() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutInvalidationContext) ContentOffsetAdjustment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutInvalidationContext) PreviousIndexPathsForInteractivelyMovingItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutInvalidationContext) TargetIndexPathsForInteractivelyMovingItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutInvalidationContext) InvalidateDecorationElementsOfKind() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutInvalidationContext) InvalidatedDecorationIndexPaths() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutInvalidationContext) ContentSizeAdjustment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutInvalidationContext) SetContentSizeAdjustment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutInvalidationContext) InvalidatedSupplementaryIndexPaths() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutInvalidationContext) SetContentOffsetAdjustment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutInvalidationContext) InvalidatedItemIndexPaths() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutInvalidationContext) InteractiveMovementTarget() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutInvalidationContext) InvalidateEverything() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutInvalidationContext) InvalidateDataSourceCounts() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

