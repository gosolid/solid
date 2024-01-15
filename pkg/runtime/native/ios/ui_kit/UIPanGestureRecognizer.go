//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIPanGestureRecognizer : UIKit.UIGestureRecognizer
*/

type UIPanGestureRecognizer struct {
  *UIGestureRecognizer

}

func (r *UIPanGestureRecognizer) SetTranslation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPanGestureRecognizer) SetMaximumNumberOfTouches() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPanGestureRecognizer) AllowedScrollTypesMask() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPanGestureRecognizer) SetAllowedScrollTypesMask() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPanGestureRecognizer) TranslationInView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPanGestureRecognizer) VelocityInView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPanGestureRecognizer) MinimumNumberOfTouches() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPanGestureRecognizer) SetMinimumNumberOfTouches() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPanGestureRecognizer) MaximumNumberOfTouches() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

