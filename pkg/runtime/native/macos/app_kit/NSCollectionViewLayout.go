//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSCollectionViewLayout : objc.NSObject
*/

type NSCollectionViewLayout struct {
  *objc.NSObject
  *foundation.NSCoding
}

func (r *NSCollectionViewLayout) InvalidateLayoutWithContext() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayout) RegisterClass() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayout) RegisterNib() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayout) CollectionView() (*NSCollectionView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayout) InvalidateLayout() error {
  return fmt.Errorf("unimplemented")
}

