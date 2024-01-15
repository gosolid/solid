//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIColorPickerViewController : UIKit.UIViewController
*/

type UIColorPickerViewController struct {
  *UIViewController

}

func (r *UIColorPickerViewController) SupportsAlpha() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIColorPickerViewController) SetSupportsAlpha() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIColorPickerViewController) InitWithNibName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIColorPickerViewController) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIColorPickerViewController) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIColorPickerViewController) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIColorPickerViewController) SelectedColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIColorPickerViewController) SetSelectedColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

