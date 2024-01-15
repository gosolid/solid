//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UITabBarItem : UIKit.UIBarItem
*/

type UITabBarItem struct {
  *UIBarItem

}

func (r *UITabBarItem) BadgeColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItem) ScrollEdgeAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItem) SetTitlePositionAdjustment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItem) InitWithTabBarSystemItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItem) SetFinishedSelectedImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItem) FinishedSelectedImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItem) FinishedUnselectedImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItem) SetBadgeTextAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItem) BadgeTextAttributesForState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItem) BadgeValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItem) SetBadgeColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItem) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItem) SelectedImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItem) SetSelectedImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItem) SetBadgeValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItem) TitlePositionAdjustment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItem) SetStandardAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItem) SetScrollEdgeAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItem) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItem) InitWithTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarItem) StandardAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

