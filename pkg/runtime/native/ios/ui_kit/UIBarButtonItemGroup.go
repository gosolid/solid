//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIBarButtonItemGroup : objc.NSObject
*/

type UIBarButtonItemGroup struct {
  *objc.NSObject

}

func (r *UIBarButtonItemGroup) MenuRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemGroup) SetHidden() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemGroup) MovableGroupWithCustomizationIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemGroup) IsDisplayingRepresentativeItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemGroup) SetAlwaysAvailable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemGroup) BarButtonItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemGroup) SetRepresentativeItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemGroup) IsHidden() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemGroup) AlwaysAvailable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemGroup) SetMenuRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemGroup) InitWithBarButtonItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemGroup) FixedGroupWithRepresentativeItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemGroup) SetBarButtonItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemGroup) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemGroup) OptionalGroupWithCustomizationIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemGroup) RepresentativeItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

