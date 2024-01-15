//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UILongPressGestureRecognizer : UIKit.UIGestureRecognizer
*/

type UILongPressGestureRecognizer struct {
  *UIGestureRecognizer

}

func (r *UILongPressGestureRecognizer) SetMinimumPressDuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILongPressGestureRecognizer) AllowableMovement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILongPressGestureRecognizer) SetAllowableMovement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILongPressGestureRecognizer) NumberOfTapsRequired() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILongPressGestureRecognizer) SetNumberOfTapsRequired() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILongPressGestureRecognizer) NumberOfTouchesRequired() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILongPressGestureRecognizer) SetNumberOfTouchesRequired() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UILongPressGestureRecognizer) MinimumPressDuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

