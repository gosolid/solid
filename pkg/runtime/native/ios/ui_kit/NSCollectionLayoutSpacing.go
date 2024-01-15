//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSCollectionLayoutSpacing : objc.NSObject
*/

type NSCollectionLayoutSpacing struct {
  *objc.NSObject

}

func (r *NSCollectionLayoutSpacing) FixedSpacing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSpacing) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSpacing) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSpacing) Spacing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSpacing) IsFlexibleSpacing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSpacing) IsFixedSpacing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSpacing) FlexibleSpacing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

