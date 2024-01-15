//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSCollectionLayoutBoundarySupplementaryItem : AppKit.NSCollectionLayoutSupplementaryItem
*/

type NSCollectionLayoutBoundarySupplementaryItem struct {
  *NSCollectionLayoutSupplementaryItem
  *foundation.NSCopying
}

func (r *NSCollectionLayoutBoundarySupplementaryItem) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutBoundarySupplementaryItem) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutBoundarySupplementaryItem) ExtendsBoundary() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutBoundarySupplementaryItem) SetPinToVisibleBounds() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutBoundarySupplementaryItem) Offset() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutBoundarySupplementaryItem) BoundarySupplementaryItemWithLayoutSize() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutBoundarySupplementaryItem) SetExtendsBoundary() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutBoundarySupplementaryItem) PinToVisibleBounds() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutBoundarySupplementaryItem) Alignment() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

