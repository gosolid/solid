//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UICollectionViewUpdateItem : objc.NSObject
*/

type UICollectionViewUpdateItem struct {
  *objc.NSObject

}

func (r *UICollectionViewUpdateItem) IndexPathBeforeUpdate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewUpdateItem) IndexPathAfterUpdate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewUpdateItem) UpdateAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

