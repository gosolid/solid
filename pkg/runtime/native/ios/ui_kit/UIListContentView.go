//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIListContentView : UIKit.UIView
*/

type UIListContentView struct {
  *UIView

}

func (r *UIListContentView) InitWithConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIListContentView) InitWithFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIListContentView) Configuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIListContentView) SetConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIListContentView) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIListContentView) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIListContentView) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIListContentView) TextLayoutGuide() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIListContentView) SecondaryTextLayoutGuide() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIListContentView) ImageLayoutGuide() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

