//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UICollectionViewLayoutAttributes : objc.NSObject
*/

type UICollectionViewLayoutAttributes struct {
  *objc.NSObject

}

func (r *UICollectionViewLayoutAttributes) Size() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) RepresentedElementCategory() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) RepresentedElementKind() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) LayoutAttributesForCellWithIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) SetHidden() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) LayoutAttributesForDecorationViewOfKind() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) SetSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) SetBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) SetAlpha() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) IsHidden() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) SetIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) LayoutAttributesForSupplementaryViewOfKind() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) Frame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) SetFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) Transform3D() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) SetZIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) Center() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) Transform() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) Alpha() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) IndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) SetTransform3D() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) Bounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) ZIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) SetCenter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewLayoutAttributes) SetTransform() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

