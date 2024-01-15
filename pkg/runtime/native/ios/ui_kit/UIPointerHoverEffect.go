//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIPointerHoverEffect : UIKit.UIPointerEffect
*/

type UIPointerHoverEffect struct {
  *UIPointerEffect

}

func (r *UIPointerHoverEffect) SetPreferredTintMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerHoverEffect) PrefersShadow() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerHoverEffect) SetPrefersShadow() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerHoverEffect) PrefersScaledContent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerHoverEffect) SetPrefersScaledContent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerHoverEffect) PreferredTintMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

