//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSCollectionLayoutItem : objc.NSObject
*/

type NSCollectionLayoutItem struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSCollectionLayoutItem) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutItem) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutItem) EdgeSpacing() (*NSCollectionLayoutEdgeSpacing, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutItem) ItemWithLayoutSize() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutItem) SetContentInsets() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutItem) SetEdgeSpacing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutItem) LayoutSize() (*NSCollectionLayoutSize, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutItem) SupplementaryItems() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutItem) ContentInsets() (NSDirectionalEdgeInsets, error) {
  return NSDirectionalEdgeInsets{}, fmt.Errorf("unimplemented")
}

