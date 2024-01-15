//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICalendarSelectionMultiDate : UIKit.UICalendarSelection
*/

type UICalendarSelectionMultiDate struct {
  *UICalendarSelection

}

func (r *UICalendarSelectionMultiDate) SetSelectedDates() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarSelectionMultiDate) InitWithDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarSelectionMultiDate) SelectedDates() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarSelectionMultiDate) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

