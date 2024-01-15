//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSCollectionLayoutSize : objc.NSObject
*/

type NSCollectionLayoutSize struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSCollectionLayoutSize) SizeWithWidthDimension() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSize) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSize) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSize) WidthDimension() (*NSCollectionLayoutDimension, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSize) HeightDimension() (*NSCollectionLayoutDimension, error) {
  return nil, fmt.Errorf("unimplemented")
}

