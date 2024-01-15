//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSCollectionViewCompositionalLayoutConfiguration : objc.NSObject
*/

type NSCollectionViewCompositionalLayoutConfiguration struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSCollectionViewCompositionalLayoutConfiguration) InterSectionSpacing() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewCompositionalLayoutConfiguration) SetInterSectionSpacing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewCompositionalLayoutConfiguration) BoundarySupplementaryItems() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewCompositionalLayoutConfiguration) SetBoundarySupplementaryItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewCompositionalLayoutConfiguration) ScrollDirection() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewCompositionalLayoutConfiguration) SetScrollDirection() error {
  return fmt.Errorf("unimplemented")
}

