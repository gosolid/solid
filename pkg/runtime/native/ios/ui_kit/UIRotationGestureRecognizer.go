//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIRotationGestureRecognizer : UIKit.UIGestureRecognizer
*/

type UIRotationGestureRecognizer struct {
  *UIGestureRecognizer

}

func (r *UIRotationGestureRecognizer) Rotation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIRotationGestureRecognizer) SetRotation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIRotationGestureRecognizer) Velocity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

