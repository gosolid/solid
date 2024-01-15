//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface AppKit.NSCollectionLayoutSpacing : objc.NSObject
*/

type NSCollectionLayoutSpacing struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSCollectionLayoutSpacing) FixedSpacing() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSpacing) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSpacing) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSpacing) Spacing() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSpacing) IsFlexibleSpacing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSpacing) IsFixedSpacing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSpacing) FlexibleSpacing() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

