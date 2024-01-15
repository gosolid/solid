//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSCollectionLayoutDimension : objc.NSObject
*/

type NSCollectionLayoutDimension struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSCollectionLayoutDimension) FractionalWidthDimension() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDimension) FractionalHeightDimension() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDimension) EstimatedDimension() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDimension) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDimension) IsFractionalHeight() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDimension) IsAbsolute() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDimension) Dimension() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDimension) AbsoluteDimension() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDimension) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDimension) IsFractionalWidth() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDimension) IsEstimated() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

