//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UITapGestureRecognizer : UIKit.UIGestureRecognizer
*/

type UITapGestureRecognizer struct {
  *UIGestureRecognizer

}

func (r *UITapGestureRecognizer) NumberOfTouchesRequired() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITapGestureRecognizer) SetNumberOfTouchesRequired() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITapGestureRecognizer) ButtonMaskRequired() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITapGestureRecognizer) SetButtonMaskRequired() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITapGestureRecognizer) NumberOfTapsRequired() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITapGestureRecognizer) SetNumberOfTapsRequired() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

