//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSCollectionLayoutSection : objc.NSObject
*/

type NSCollectionLayoutSection struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSCollectionLayoutSection) InterGroupSpacing() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSection) SetInterGroupSpacing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSection) OrthogonalScrollingBehavior() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSection) SupplementariesFollowContentInsets() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSection) SetSupplementariesFollowContentInsets() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSection) SetVisibleItemsInvalidationHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSection) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSection) SetContentInsets() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSection) DecorationItems() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSection) SectionWithGroup() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSection) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSection) VisibleItemsInvalidationHandler() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSection) SetDecorationItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSection) ContentInsets() (NSDirectionalEdgeInsets, error) {
  return NSDirectionalEdgeInsets{}, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSection) BoundarySupplementaryItems() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSection) SetOrthogonalScrollingBehavior() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutSection) SetBoundarySupplementaryItems() error {
  return fmt.Errorf("unimplemented")
}

