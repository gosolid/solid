//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSCollectionLayoutEdgeSpacing : objc.NSObject
*/

type NSCollectionLayoutEdgeSpacing struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSCollectionLayoutEdgeSpacing) SpacingForLeading() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutEdgeSpacing) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutEdgeSpacing) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutEdgeSpacing) Leading() (*NSCollectionLayoutSpacing, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutEdgeSpacing) Top() (*NSCollectionLayoutSpacing, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutEdgeSpacing) Trailing() (*NSCollectionLayoutSpacing, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutEdgeSpacing) Bottom() (*NSCollectionLayoutSpacing, error) {
  return nil, fmt.Errorf("unimplemented")
}

