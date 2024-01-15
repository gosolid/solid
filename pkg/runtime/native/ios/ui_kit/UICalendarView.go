//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICalendarView : UIKit.UIView
*/

type UICalendarView struct {
  *UIView

}

func (r *UICalendarView) WantsDateDecorations() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarView) ReloadDecorationsForDateComponents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarView) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarView) SetLocale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarView) SetTimeZone() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarView) SetAvailableDateRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarView) AvailableDateRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarView) SetVisibleDateComponents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarView) SelectionBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarView) Locale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarView) TimeZone() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarView) FontDesign() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarView) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarView) Calendar() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarView) SetCalendar() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarView) SetFontDesign() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarView) VisibleDateComponents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarView) SetSelectionBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICalendarView) SetWantsDateDecorations() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

