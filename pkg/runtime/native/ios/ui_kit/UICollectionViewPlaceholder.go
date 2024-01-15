//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UICollectionViewPlaceholder : objc.NSObject
*/

type UICollectionViewPlaceholder struct {
  *objc.NSObject

}

func (r *UICollectionViewPlaceholder) InitWithInsertionIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewPlaceholder) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewPlaceholder) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewPlaceholder) CellUpdateHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewPlaceholder) SetCellUpdateHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

