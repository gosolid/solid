//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSComboBoxCellDataSource
*/

type NSComboBoxCellDataSource struct {

}

func (r *NSComboBoxCellDataSource) NumberOfItemsInComboBoxCell() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSComboBoxCellDataSource) ComboBoxCell() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

