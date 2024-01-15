//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSCollectionViewLayoutAttributes : objc.NSObject
*/

type NSCollectionViewLayoutAttributes struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSCollectionViewLayoutAttributes) LayoutAttributesForItemWithIndexPath() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutAttributes) LayoutAttributesForInterItemGapBeforeIndexPath() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutAttributes) LayoutAttributesForSupplementaryViewOfKind() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutAttributes) Size() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutAttributes) SetSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutAttributes) IsHidden() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutAttributes) LayoutAttributesForDecorationViewOfKind() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutAttributes) ZIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutAttributes) SetIndexPath() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutAttributes) Frame() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutAttributes) SetZIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutAttributes) SetHidden() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutAttributes) IndexPath() (*foundation.NSIndexPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutAttributes) SetFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutAttributes) Alpha() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutAttributes) SetAlpha() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutAttributes) RepresentedElementCategory() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewLayoutAttributes) RepresentedElementKind() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

