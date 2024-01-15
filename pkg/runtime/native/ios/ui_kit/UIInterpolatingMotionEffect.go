//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIInterpolatingMotionEffect : UIKit.UIMotionEffect
*/

type UIInterpolatingMotionEffect struct {
  *UIMotionEffect

}

func (r *UIInterpolatingMotionEffect) InitWithKeyPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIInterpolatingMotionEffect) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIInterpolatingMotionEffect) KeyPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIInterpolatingMotionEffect) Type() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIInterpolatingMotionEffect) MinimumRelativeValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIInterpolatingMotionEffect) SetMinimumRelativeValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIInterpolatingMotionEffect) MaximumRelativeValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIInterpolatingMotionEffect) SetMaximumRelativeValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

