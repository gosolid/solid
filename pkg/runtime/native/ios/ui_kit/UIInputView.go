//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIInputView : UIKit.UIView
*/

type UIInputView struct {
  *UIView

}

func (r *UIInputView) AllowsSelfSizing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIInputView) SetAllowsSelfSizing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIInputView) InitWithFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIInputView) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIInputView) InputViewStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

