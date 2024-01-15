//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICalendarSelectionSingleDate : UIKit.UICalendarSelection
*/

type UICalendarSelectionSingleDate struct {
  *UICalendarSelection

}

func (r *UICalendarSelectionSingleDate) SelectedDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarSelectionSingleDate) SetSelectedDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarSelectionSingleDate) InitWithDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarSelectionSingleDate) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

