//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIActivityIndicatorView : UIKit.UIView
*/

type UIActivityIndicatorView struct {
  *UIView

}

func (r *UIActivityIndicatorView) StopAnimating() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityIndicatorView) HidesWhenStopped() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityIndicatorView) Color() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityIndicatorView) StartAnimating() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityIndicatorView) InitWithFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityIndicatorView) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityIndicatorView) ActivityIndicatorViewStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityIndicatorView) SetActivityIndicatorViewStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityIndicatorView) SetHidesWhenStopped() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityIndicatorView) SetColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityIndicatorView) IsAnimating() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIActivityIndicatorView) InitWithActivityIndicatorStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

