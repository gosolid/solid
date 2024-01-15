//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UICollectionViewDiffableDataSourceReorderingHandlers : objc.NSObject
*/

type UICollectionViewDiffableDataSourceReorderingHandlers struct {
  *objc.NSObject

}

func (r *UICollectionViewDiffableDataSourceReorderingHandlers) SetWillReorderHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSourceReorderingHandlers) DidReorderHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSourceReorderingHandlers) SetDidReorderHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSourceReorderingHandlers) CanReorderItemHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSourceReorderingHandlers) SetCanReorderItemHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSourceReorderingHandlers) WillReorderHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

