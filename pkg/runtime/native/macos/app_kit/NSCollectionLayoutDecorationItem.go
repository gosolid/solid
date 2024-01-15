//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSCollectionLayoutDecorationItem : AppKit.NSCollectionLayoutItem
*/

type NSCollectionLayoutDecorationItem struct {
  *NSCollectionLayoutItem
  *foundation.NSCopying
}

func (r *NSCollectionLayoutDecorationItem) ZIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDecorationItem) SetZIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDecorationItem) ElementKind() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDecorationItem) BackgroundDecorationItemWithElementKind() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDecorationItem) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDecorationItem) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

