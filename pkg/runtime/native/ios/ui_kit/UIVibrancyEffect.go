//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIVibrancyEffect : UIKit.UIVisualEffect
*/

type UIVibrancyEffect struct {
  *UIVisualEffect

}

func (r *UIVibrancyEffect) EffectForBlurEffect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

