//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UICubicTimingParameters : objc.NSObject
*/

type UICubicTimingParameters struct {
  *objc.NSObject

}

func (r *UICubicTimingParameters) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICubicTimingParameters) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICubicTimingParameters) InitWithAnimationCurve() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICubicTimingParameters) InitWithControlPoint1() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICubicTimingParameters) AnimationCurve() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICubicTimingParameters) ControlPoint1() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICubicTimingParameters) ControlPoint2() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

