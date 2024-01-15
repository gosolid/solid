//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UITabBarItemStateAppearance : objc.NSObject
*/

type UITabBarItemStateAppearance struct {
  *objc.NSObject

}

func (r *UITabBarItemStateAppearance) SetBadgeTextAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItemStateAppearance) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItemStateAppearance) SetTitleTextAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItemStateAppearance) SetTitlePositionAdjustment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItemStateAppearance) SetBadgePositionAdjustment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItemStateAppearance) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItemStateAppearance) TitleTextAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItemStateAppearance) SetIconColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItemStateAppearance) BadgeTextAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItemStateAppearance) BadgePositionAdjustment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItemStateAppearance) SetBadgeBackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItemStateAppearance) BadgeTitlePositionAdjustment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItemStateAppearance) TitlePositionAdjustment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItemStateAppearance) IconColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItemStateAppearance) BadgeBackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItemStateAppearance) SetBadgeTitlePositionAdjustment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

