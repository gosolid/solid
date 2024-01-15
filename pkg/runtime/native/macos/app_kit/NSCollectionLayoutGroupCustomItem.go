//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface AppKit.NSCollectionLayoutGroupCustomItem : objc.NSObject
*/

type NSCollectionLayoutGroupCustomItem struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSCollectionLayoutGroupCustomItem) CustomItemWithFrame() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroupCustomItem) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroupCustomItem) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroupCustomItem) Frame() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutGroupCustomItem) ZIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

