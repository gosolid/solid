//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSDatePickerCellDelegate
*/

type NSDatePickerCellDelegate struct {

}

func (r *NSDatePickerCellDelegate) DatePickerCell() error {
  return fmt.Errorf("unimplemented")
}

