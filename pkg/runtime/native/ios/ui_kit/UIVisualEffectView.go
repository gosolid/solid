//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIVisualEffectView : UIKit.UIView
*/

type UIVisualEffectView struct {
  *UIView

}

func (r *UIVisualEffectView) ContentView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIVisualEffectView) Effect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIVisualEffectView) SetEffect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIVisualEffectView) InitWithEffect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIVisualEffectView) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

