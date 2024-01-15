//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UITabBarAppearance : UIKit.UIBarAppearance
*/

type UITabBarAppearance struct {
  *UIBarAppearance

}

func (r *UITabBarAppearance) SetStackedLayoutAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarAppearance) CompactInlineLayoutAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarAppearance) SelectionIndicatorTintColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarAppearance) SelectionIndicatorImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarAppearance) StackedItemWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarAppearance) StackedLayoutAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarAppearance) InlineLayoutAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarAppearance) SetSelectionIndicatorTintColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarAppearance) SetStackedItemWidth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarAppearance) SetStackedItemSpacing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarAppearance) SetCompactInlineLayoutAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarAppearance) StackedItemPositioning() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarAppearance) SetInlineLayoutAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarAppearance) SetSelectionIndicatorImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarAppearance) SetStackedItemPositioning() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITabBarAppearance) StackedItemSpacing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

