//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.NSCollectionLayoutGroup : UIKit.NSCollectionLayoutItem
*/

type NSCollectionLayoutGroup struct {
  *NSCollectionLayoutItem

}

func (r *NSCollectionLayoutGroup) VerticalGroupWithLayoutSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroup) CustomGroupWithLayoutSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroup) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroup) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroup) VisualDescription() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroup) SetSupplementaryItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroup) InterItemSpacing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroup) HorizontalGroupWithLayoutSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroup) SupplementaryItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroup) SetInterItemSpacing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroup) Subitems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

