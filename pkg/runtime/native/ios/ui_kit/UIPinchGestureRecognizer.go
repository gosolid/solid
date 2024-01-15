//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIPinchGestureRecognizer : UIKit.UIGestureRecognizer
*/

type UIPinchGestureRecognizer struct {
  *UIGestureRecognizer

}

func (r *UIPinchGestureRecognizer) Scale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPinchGestureRecognizer) SetScale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPinchGestureRecognizer) Velocity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

