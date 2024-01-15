//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface AppKit.NSCollectionViewUpdateItem : objc.NSObject
*/

type NSCollectionViewUpdateItem struct {
  *objc.NSObject

}

func (r *NSCollectionViewUpdateItem) IndexPathBeforeUpdate() (*foundation.NSIndexPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewUpdateItem) IndexPathAfterUpdate() (*foundation.NSIndexPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewUpdateItem) UpdateAction() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

