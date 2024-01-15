//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSCollectionLayoutSupplementaryItem : AppKit.NSCollectionLayoutItem
*/

type NSCollectionLayoutSupplementaryItem struct {
  *NSCollectionLayoutItem
  *foundation.NSCopying
}

func (r *NSCollectionLayoutSupplementaryItem) ZIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSupplementaryItem) SetZIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSupplementaryItem) ElementKind() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSupplementaryItem) ContainerAnchor() (*NSCollectionLayoutAnchor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSupplementaryItem) ItemAnchor() (*NSCollectionLayoutAnchor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSupplementaryItem) SupplementaryItemWithLayoutSize() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSupplementaryItem) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSupplementaryItem) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

