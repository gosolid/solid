//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSCollectionLayoutGroupCustomItem : objc.NSObject
*/

type NSCollectionLayoutGroupCustomItem struct {
  *objc.NSObject

}

func (r *NSCollectionLayoutGroupCustomItem) CustomItemWithFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroupCustomItem) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroupCustomItem) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroupCustomItem) Frame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroupCustomItem) ZIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

