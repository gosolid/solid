//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UISpringTimingParameters : objc.NSObject
*/

type UISpringTimingParameters struct {
  *objc.NSObject

}

func (r *UISpringTimingParameters) InitWithMass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISpringTimingParameters) InitWithDuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISpringTimingParameters) InitialVelocity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISpringTimingParameters) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISpringTimingParameters) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISpringTimingParameters) InitWithDampingRatio() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

