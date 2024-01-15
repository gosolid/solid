//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSComboBoxDelegate
*/

type NSComboBoxDelegate struct {

}

func (r *NSComboBoxDelegate) ComboBoxWillDismiss() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBoxDelegate) ComboBoxSelectionDidChange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBoxDelegate) ComboBoxSelectionIsChanging() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSComboBoxDelegate) ComboBoxWillPopUp() error {
  return fmt.Errorf("unimplemented")
}

