//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIViewPropertyAnimator : objc.NSObject
*/

type UIViewPropertyAnimator struct {
  *objc.NSObject

}

func (r *UIViewPropertyAnimator) InitWithDuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewPropertyAnimator) AddCompletion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewPropertyAnimator) Delay() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewPropertyAnimator) IsManualHitTestingEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewPropertyAnimator) IsInterruptible() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewPropertyAnimator) SetScrubsLinearly() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewPropertyAnimator) PausesOnCompletion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewPropertyAnimator) RunningPropertyAnimatorWithDuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewPropertyAnimator) AddAnimations() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewPropertyAnimator) ContinueAnimationWithTimingParameters() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewPropertyAnimator) Duration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewPropertyAnimator) IsUserInteractionEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewPropertyAnimator) SetUserInteractionEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewPropertyAnimator) ScrubsLinearly() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewPropertyAnimator) TimingParameters() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewPropertyAnimator) SetManualHitTestingEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewPropertyAnimator) SetInterruptible() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIViewPropertyAnimator) SetPausesOnCompletion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

