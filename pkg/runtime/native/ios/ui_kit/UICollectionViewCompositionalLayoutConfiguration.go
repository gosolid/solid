//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UICollectionViewCompositionalLayoutConfiguration : objc.NSObject
*/

type UICollectionViewCompositionalLayoutConfiguration struct {
  *objc.NSObject

}

func (r *UICollectionViewCompositionalLayoutConfiguration) SetScrollDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCompositionalLayoutConfiguration) InterSectionSpacing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCompositionalLayoutConfiguration) SetInterSectionSpacing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCompositionalLayoutConfiguration) BoundarySupplementaryItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCompositionalLayoutConfiguration) SetBoundarySupplementaryItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCompositionalLayoutConfiguration) ContentInsetsReference() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCompositionalLayoutConfiguration) SetContentInsetsReference() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewCompositionalLayoutConfiguration) ScrollDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

