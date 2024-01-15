//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UINavigationBarAppearance : UIKit.UIBarAppearance
*/

type UINavigationBarAppearance struct {
  *UIBarAppearance

}

func (r *UINavigationBarAppearance) SetTitlePositionAdjustment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationBarAppearance) SetDoneButtonAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationBarAppearance) SetBackButtonAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationBarAppearance) TitlePositionAdjustment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationBarAppearance) DoneButtonAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationBarAppearance) TitleTextAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationBarAppearance) SetTitleTextAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationBarAppearance) SetLargeTitleTextAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationBarAppearance) ButtonAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationBarAppearance) BackButtonAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationBarAppearance) SetBackIndicatorImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationBarAppearance) LargeTitleTextAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationBarAppearance) SetButtonAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationBarAppearance) BackIndicatorImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINavigationBarAppearance) BackIndicatorTransitionMaskImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

