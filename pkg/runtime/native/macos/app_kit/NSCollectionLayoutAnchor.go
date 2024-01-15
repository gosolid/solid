//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSCollectionLayoutAnchor : objc.NSObject
*/

type NSCollectionLayoutAnchor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSCollectionLayoutAnchor) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutAnchor) Edges() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutAnchor) Offset() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutAnchor) IsAbsoluteOffset() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutAnchor) IsFractionalOffset() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutAnchor) LayoutAnchorWithEdges() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutAnchor) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

