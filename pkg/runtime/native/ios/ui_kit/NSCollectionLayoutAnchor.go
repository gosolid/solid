//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSCollectionLayoutAnchor : objc.NSObject
*/

type NSCollectionLayoutAnchor struct {
  *objc.NSObject

}

func (r *NSCollectionLayoutAnchor) LayoutAnchorWithEdges() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutAnchor) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutAnchor) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutAnchor) Edges() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutAnchor) Offset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutAnchor) IsAbsoluteOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutAnchor) IsFractionalOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

