//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIToolbarAppearance : UIKit.UIBarAppearance
*/

type UIToolbarAppearance struct {
  *UIBarAppearance

}

func (r *UIToolbarAppearance) DoneButtonAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIToolbarAppearance) SetDoneButtonAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIToolbarAppearance) ButtonAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIToolbarAppearance) SetButtonAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

