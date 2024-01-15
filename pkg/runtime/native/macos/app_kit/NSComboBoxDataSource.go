//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSComboBoxDataSource
*/

type NSComboBoxDataSource struct {

}

func (r *NSComboBoxDataSource) NumberOfItemsInComboBox() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSComboBoxDataSource) ComboBox() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

