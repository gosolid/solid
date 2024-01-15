//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIFocusHaloEffect : UIKit.UIFocusEffect
*/

type UIFocusHaloEffect struct {
  *UIFocusEffect

}

func (r *UIFocusHaloEffect) SetContainerView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusHaloEffect) Position() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusHaloEffect) SetPosition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusHaloEffect) EffectWithRoundedRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusHaloEffect) EffectWithPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusHaloEffect) ContainerView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusHaloEffect) EffectWithRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusHaloEffect) ReferenceView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusHaloEffect) SetReferenceView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

