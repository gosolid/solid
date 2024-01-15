//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIKeyboardLayoutGuide : UIKit.UITrackingLayoutGuide
*/

type UIKeyboardLayoutGuide struct {
  *UITrackingLayoutGuide

}

func (r *UIKeyboardLayoutGuide) FollowsUndockedKeyboard() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyboardLayoutGuide) SetFollowsUndockedKeyboard() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyboardLayoutGuide) UsesBottomSafeArea() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyboardLayoutGuide) SetUsesBottomSafeArea() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyboardLayoutGuide) KeyboardDismissPadding() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyboardLayoutGuide) SetKeyboardDismissPadding() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

