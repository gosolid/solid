//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.NSCollectionLayoutDecorationItem : UIKit.NSCollectionLayoutItem
*/

type NSCollectionLayoutDecorationItem struct {
  *NSCollectionLayoutItem

}

func (r *NSCollectionLayoutDecorationItem) ElementKind() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDecorationItem) BackgroundDecorationItemWithElementKind() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDecorationItem) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDecorationItem) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDecorationItem) ZIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDecorationItem) SetZIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

