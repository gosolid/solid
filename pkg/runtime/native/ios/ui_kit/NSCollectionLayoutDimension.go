//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSCollectionLayoutDimension : objc.NSObject
*/

type NSCollectionLayoutDimension struct {
  *objc.NSObject

}

func (r *NSCollectionLayoutDimension) FractionalWidthDimension() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDimension) FractionalHeightDimension() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDimension) AbsoluteDimension() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDimension) UniformAcrossSiblingsWithEstimate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDimension) IsEstimated() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDimension) Dimension() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDimension) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDimension) IsFractionalHeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDimension) IsAbsolute() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDimension) EstimatedDimension() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDimension) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDimension) IsFractionalWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutDimension) IsUniformAcrossSiblings() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

