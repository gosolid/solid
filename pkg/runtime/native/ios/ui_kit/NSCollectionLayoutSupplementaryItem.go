//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.NSCollectionLayoutSupplementaryItem : UIKit.NSCollectionLayoutItem
*/

type NSCollectionLayoutSupplementaryItem struct {
  *NSCollectionLayoutItem

}

func (r *NSCollectionLayoutSupplementaryItem) SetZIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSupplementaryItem) ElementKind() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSupplementaryItem) ContainerAnchor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSupplementaryItem) ItemAnchor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSupplementaryItem) SupplementaryItemWithLayoutSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSupplementaryItem) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSupplementaryItem) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSupplementaryItem) ZIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

