//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPercentDrivenInteractiveTransition : objc.NSObject
*/

type UIPercentDrivenInteractiveTransition struct {
  *objc.NSObject

}

func (r *UIPercentDrivenInteractiveTransition) Duration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPercentDrivenInteractiveTransition) SetCompletionSpeed() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPercentDrivenInteractiveTransition) TimingCurve() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPercentDrivenInteractiveTransition) WantsInteractiveStart() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPercentDrivenInteractiveTransition) CancelInteractiveTransition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPercentDrivenInteractiveTransition) SetCompletionCurve() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPercentDrivenInteractiveTransition) SetWantsInteractiveStart() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPercentDrivenInteractiveTransition) PauseInteractiveTransition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPercentDrivenInteractiveTransition) UpdateInteractiveTransition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPercentDrivenInteractiveTransition) FinishInteractiveTransition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPercentDrivenInteractiveTransition) PercentComplete() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPercentDrivenInteractiveTransition) CompletionSpeed() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPercentDrivenInteractiveTransition) CompletionCurve() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPercentDrivenInteractiveTransition) SetTimingCurve() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

