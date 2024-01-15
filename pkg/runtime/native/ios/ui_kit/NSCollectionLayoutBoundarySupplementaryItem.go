//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.NSCollectionLayoutBoundarySupplementaryItem : UIKit.NSCollectionLayoutSupplementaryItem
*/

type NSCollectionLayoutBoundarySupplementaryItem struct {
  *NSCollectionLayoutSupplementaryItem

}

func (r *NSCollectionLayoutBoundarySupplementaryItem) BoundarySupplementaryItemWithLayoutSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutBoundarySupplementaryItem) Offset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutBoundarySupplementaryItem) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutBoundarySupplementaryItem) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutBoundarySupplementaryItem) ExtendsBoundary() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutBoundarySupplementaryItem) SetExtendsBoundary() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutBoundarySupplementaryItem) PinToVisibleBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutBoundarySupplementaryItem) SetPinToVisibleBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutBoundarySupplementaryItem) Alignment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

