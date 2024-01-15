//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UIGestureRecognizer : objc.NSObject
*/

type UIGestureRecognizer struct {
  *objc.NSObject

}

func (r *UIGestureRecognizer) IsEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) AllowedTouchTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) ModifierFlags() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) SetAllowedTouchTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) State() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) RequireGestureRecognizerToFail() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) SetDelaysTouchesEnded() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) RemoveTarget() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) SetRequiresExclusiveTouchType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) NumberOfTouches() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) LocationInView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) AllowedPressTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) LocationOfTouch() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) DelaysTouchesEnded() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) SetName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) InitWithTarget() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) CancelsTouchesInView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) SetCancelsTouchesInView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) DelaysTouchesBegan() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) ButtonMask() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) AddTarget() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) View() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) SetDelaysTouchesBegan() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) SetAllowedPressTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) RequiresExclusiveTouchType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIGestureRecognizer) SetEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

