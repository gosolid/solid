//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIRefreshControl : UIKit.UIControl
*/

type UIRefreshControl struct {
  *UIControl

}

func (r *UIRefreshControl) EndRefreshing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIRefreshControl) IsRefreshing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIRefreshControl) TintColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIRefreshControl) SetTintColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIRefreshControl) AttributedTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIRefreshControl) SetAttributedTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIRefreshControl) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIRefreshControl) BeginRefreshing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

