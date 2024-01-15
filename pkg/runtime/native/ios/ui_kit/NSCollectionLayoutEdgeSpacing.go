//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.NSCollectionLayoutEdgeSpacing : objc.NSObject
*/

type NSCollectionLayoutEdgeSpacing struct {
  *objc.NSObject

}

func (r *NSCollectionLayoutEdgeSpacing) Bottom() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutEdgeSpacing) SpacingForLeading() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutEdgeSpacing) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutEdgeSpacing) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutEdgeSpacing) Leading() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutEdgeSpacing) Top() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutEdgeSpacing) Trailing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

