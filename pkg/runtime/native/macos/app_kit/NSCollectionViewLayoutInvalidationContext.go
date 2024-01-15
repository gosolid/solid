//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSCollectionViewLayoutInvalidationContext : objc.NSObject
*/

type NSCollectionViewLayoutInvalidationContext struct {
  *objc.NSObject

}

func (r *NSCollectionViewLayoutInvalidationContext) ContentSizeAdjustment() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutInvalidationContext) SetContentSizeAdjustment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutInvalidationContext) InvalidatedSupplementaryIndexPaths() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutInvalidationContext) InvalidatedDecorationIndexPaths() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutInvalidationContext) ContentOffsetAdjustment() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutInvalidationContext) SetContentOffsetAdjustment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutInvalidationContext) InvalidateDataSourceCounts() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutInvalidationContext) InvalidatedItemIndexPaths() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutInvalidationContext) InvalidateItemsAtIndexPaths() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutInvalidationContext) InvalidateSupplementaryElementsOfKind() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutInvalidationContext) InvalidateDecorationElementsOfKind() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutInvalidationContext) InvalidateEverything() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

