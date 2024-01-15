//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSCollectionLayoutGroup : AppKit.NSCollectionLayoutItem
*/

type NSCollectionLayoutGroup struct {
  *NSCollectionLayoutItem
  *foundation.NSCopying
}

func (r *NSCollectionLayoutGroup) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroup) SupplementaryItems() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroup) SetSupplementaryItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroup) SetInterItemSpacing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroup) VerticalGroupWithLayoutSize() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroup) CustomGroupWithLayoutSize() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroup) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroup) VisualDescription() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroup) InterItemSpacing() (*NSCollectionLayoutSpacing, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroup) Subitems() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroup) HorizontalGroupWithLayoutSize() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

