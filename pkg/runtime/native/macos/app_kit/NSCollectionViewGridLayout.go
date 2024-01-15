//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSCollectionViewGridLayout : AppKit.NSCollectionViewLayout
*/

type NSCollectionViewGridLayout struct {
  *NSCollectionViewLayout

}

func (r *NSCollectionViewGridLayout) Margins() (foundation.NSEdgeInsets, error) {
  return foundation.NSEdgeInsets{}, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewGridLayout) SetMinimumLineSpacing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewGridLayout) MaximumNumberOfColumns() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewGridLayout) MaximumItemSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewGridLayout) SetBackgroundColors() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewGridLayout) SetMaximumNumberOfRows() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewGridLayout) SetMaximumNumberOfColumns() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewGridLayout) MinimumItemSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewGridLayout) SetMinimumItemSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewGridLayout) BackgroundColors() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewGridLayout) MaximumNumberOfRows() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewGridLayout) SetMaximumItemSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewGridLayout) SetMargins() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewGridLayout) MinimumInteritemSpacing() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewGridLayout) SetMinimumInteritemSpacing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewGridLayout) MinimumLineSpacing() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

