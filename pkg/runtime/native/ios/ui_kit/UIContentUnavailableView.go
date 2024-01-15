//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIContentUnavailableView : UIKit.UIView
*/

type UIContentUnavailableView struct {
  *UIView

}

func (r *UIContentUnavailableView) Configuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableView) SetConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableView) SetScrollEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableView) InitWithConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableView) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableView) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableView) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableView) InitWithFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableView) IsScrollEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

