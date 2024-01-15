//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIBlurEffect : UIKit.UIVisualEffect
*/

type UIBlurEffect struct {
  *UIVisualEffect

}

func (r *UIBlurEffect) EffectWithStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

