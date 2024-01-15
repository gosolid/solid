//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIPopoverBackgroundView : UIKit.UIView
*/

type UIPopoverBackgroundView struct {
  *UIView

}

func (r *UIPopoverBackgroundView) ArrowOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPopoverBackgroundView) SetArrowOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPopoverBackgroundView) ArrowDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPopoverBackgroundView) SetArrowDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPopoverBackgroundView) WantsDefaultContentAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

