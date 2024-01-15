//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIHoverGestureRecognizer : UIKit.UIGestureRecognizer
*/

type UIHoverGestureRecognizer struct {
  *UIGestureRecognizer

}

func (r *UIHoverGestureRecognizer) AzimuthAngleInView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIHoverGestureRecognizer) AzimuthUnitVectorInView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIHoverGestureRecognizer) ZOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIHoverGestureRecognizer) AltitudeAngle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

