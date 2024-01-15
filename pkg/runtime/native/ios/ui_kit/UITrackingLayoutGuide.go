//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UITrackingLayoutGuide : UIKit.UILayoutGuide
*/

type UITrackingLayoutGuide struct {
  *UILayoutGuide

}

func (r *UITrackingLayoutGuide) RemoveAllTrackedConstraints() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITrackingLayoutGuide) SetConstraints() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITrackingLayoutGuide) ConstraintsActiveWhenNearEdge() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITrackingLayoutGuide) ConstraintsActiveWhenAwayFromEdge() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

