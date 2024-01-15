//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIEvent : objc.NSObject
*/

type UIEvent struct {
  *objc.NSObject

}

func (r *UIEvent) CoalescedTouchesForTouch() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEvent) PredictedTouchesForTouch() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEvent) Type() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEvent) ButtonMask() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEvent) AllTouches() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEvent) TouchesForWindow() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEvent) TouchesForView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEvent) TouchesForGestureRecognizer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEvent) Subtype() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEvent) Timestamp() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEvent) ModifierFlags() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

