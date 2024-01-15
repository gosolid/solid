//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol AppKit.NSCollectionLayoutVisibleItem
*/

type NSCollectionLayoutVisibleItem struct {

}

func (r *NSCollectionLayoutVisibleItem) Alpha() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutVisibleItem) ZIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutVisibleItem) SetZIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutVisibleItem) IsHidden() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutVisibleItem) SetHidden() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutVisibleItem) IndexPath() (*foundation.NSIndexPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutVisibleItem) Frame() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutVisibleItem) RepresentedElementKind() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutVisibleItem) SetAlpha() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutVisibleItem) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutVisibleItem) Bounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutVisibleItem) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutVisibleItem) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionLayoutVisibleItem) RepresentedElementCategory() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

