//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIFontPickerViewController : UIKit.UIViewController
*/

type UIFontPickerViewController struct {
  *UIViewController

}

func (r *UIFontPickerViewController) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontPickerViewController) SelectedFontDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontPickerViewController) SetSelectedFontDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontPickerViewController) InitWithConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontPickerViewController) InitWithNibName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontPickerViewController) Configuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFontPickerViewController) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

