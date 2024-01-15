//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIFocusGuide : UIKit.UILayoutGuide
*/

type UIFocusGuide struct {
  *UILayoutGuide

}

func (r *UIFocusGuide) SetEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusGuide) PreferredFocusEnvironments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusGuide) SetPreferredFocusEnvironments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusGuide) PreferredFocusedView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusGuide) SetPreferredFocusedView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusGuide) IsEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

