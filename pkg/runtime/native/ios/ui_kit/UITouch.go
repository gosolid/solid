//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UITouch : objc.NSObject
*/

type UITouch struct {
  *objc.NSObject

}

func (r *UITouch) PrecisePreviousLocationInView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITouch) AzimuthUnitVectorInView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITouch) MajorRadius() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITouch) Force() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITouch) Timestamp() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITouch) Phase() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITouch) TapCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITouch) Type() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITouch) AltitudeAngle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITouch) EstimatedPropertiesExpectingUpdates() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITouch) PreviousLocationInView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITouch) PreciseLocationInView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITouch) AzimuthAngleInView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITouch) EstimatedProperties() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITouch) LocationInView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITouch) MajorRadiusTolerance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITouch) Window() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITouch) View() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITouch) GestureRecognizers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITouch) MaximumPossibleForce() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITouch) EstimationUpdateIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

