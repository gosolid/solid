//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UISwipeGestureRecognizer : UIKit.UIGestureRecognizer
*/

type UISwipeGestureRecognizer struct {
  *UIGestureRecognizer

}

func (r *UISwipeGestureRecognizer) NumberOfTouchesRequired() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISwipeGestureRecognizer) SetNumberOfTouchesRequired() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISwipeGestureRecognizer) Direction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISwipeGestureRecognizer) SetDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

