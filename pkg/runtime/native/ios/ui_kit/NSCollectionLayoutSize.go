//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSCollectionLayoutSize : objc.NSObject
*/

type NSCollectionLayoutSize struct {
  *objc.NSObject

}

func (r *NSCollectionLayoutSize) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSize) WidthDimension() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSize) HeightDimension() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSize) SizeWithWidthDimension() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSize) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

