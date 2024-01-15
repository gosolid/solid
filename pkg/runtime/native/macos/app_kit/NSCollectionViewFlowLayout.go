//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSCollectionViewFlowLayout : AppKit.NSCollectionViewLayout
*/

type NSCollectionViewFlowLayout struct {
  *NSCollectionViewLayout

}

func (r *NSCollectionViewFlowLayout) SetEstimatedItemSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayout) SetSectionInset() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayout) SectionHeadersPinToVisibleBounds() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayout) SectionAtIndexIsCollapsed() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayout) ScrollDirection() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayout) HeaderReferenceSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayout) SectionInset() (foundation.NSEdgeInsets, error) {
  return foundation.NSEdgeInsets{}, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayout) SectionFootersPinToVisibleBounds() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayout) MinimumLineSpacing() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayout) FooterReferenceSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayout) SetFooterReferenceSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayout) SetSectionHeadersPinToVisibleBounds() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayout) ExpandSectionAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayout) SetMinimumLineSpacing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayout) MinimumInteritemSpacing() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayout) SetMinimumInteritemSpacing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayout) ItemSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayout) SetItemSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayout) EstimatedItemSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayout) SetScrollDirection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayout) CollapseSectionAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayout) SetSectionFootersPinToVisibleBounds() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewFlowLayout) SetHeaderReferenceSize() error {
  return fmt.Errorf("unimplemented")
}

