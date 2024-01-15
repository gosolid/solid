//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIMotionEffectGroup : UIKit.UIMotionEffect
*/

type UIMotionEffectGroup struct {
  *UIMotionEffect

}

func (r *UIMotionEffectGroup) MotionEffects() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMotionEffectGroup) SetMotionEffects() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

