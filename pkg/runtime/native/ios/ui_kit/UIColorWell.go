//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIColorWell : UIKit.UIControl
*/

type UIColorWell struct {
  *UIControl

}

func (r *UIColorWell) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIColorWell) SetTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIColorWell) SupportsAlpha() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIColorWell) SetSupportsAlpha() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIColorWell) SelectedColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIColorWell) SetSelectedColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

