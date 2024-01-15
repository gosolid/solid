//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UITraitCollection : objc.NSObject
*/

type UITraitCollection struct {
  *objc.NSObject

}

func (r *UITraitCollection) LayoutDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) HorizontalSizeClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) TraitCollectionWithTraitsFromCollections() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) TraitCollectionWithAccessibilityContrast() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) UserInterfaceStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) TraitCollectionWithUserInterfaceLevel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) AccessibilityContrast() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) TraitCollectionWithForceTouchCapability() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) TraitCollectionWithTypesettingLanguage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) DisplayGamut() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) TypesettingLanguage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) ContainsTraitsInCollection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) TraitCollectionWithDisplayScale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) TraitCollectionWithDisplayGamut() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) UserInterfaceLevel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) LegibilityWeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) ToolbarItemPresentationSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) ImageDynamicRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) TraitCollectionWithUserInterfaceIdiom() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) TraitCollectionWithVerticalSizeClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) TraitCollectionWithSceneCaptureState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) UserInterfaceIdiom() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) SceneCaptureState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) TraitCollectionWithUserInterfaceStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) TraitCollectionWithPreferredContentSizeCategory() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) TraitCollectionWithLegibilityWeight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) TraitCollectionWithToolbarItemPresentationSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) TraitCollectionWithImageDynamicRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) VerticalSizeClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) PreferredContentSizeCategory() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) ActiveAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) TraitCollectionWithLayoutDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) TraitCollectionWithHorizontalSizeClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) TraitCollectionWithActiveAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) DisplayScale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITraitCollection) ForceTouchCapability() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

