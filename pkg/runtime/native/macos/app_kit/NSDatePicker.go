//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSDatePicker : AppKit.NSControl
*/

type NSDatePicker struct {
  *NSControl

}

func (r *NSDatePicker) IsBordered() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) TextColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) SetDatePickerMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) SetDateValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) SetMaxDate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) DrawsBackground() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) DatePickerElements() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) SetCalendar() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) SetLocale() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) SetMinDate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) SetDatePickerStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) SetBezeled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) SetDrawsBackground() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) MinDate() (*foundation.NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) DatePickerStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) BackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) Locale() (*foundation.NSLocale, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) SetBackgroundColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) SetDatePickerElements() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) Calendar() (*foundation.NSCalendar, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) DateValue() (*foundation.NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) MaxDate() (*foundation.NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) SetPresentsCalendarOverlay() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) SetBordered() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) SetTextColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) SetTimeZone() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) SetTimeInterval() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) PresentsCalendarOverlay() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) IsBezeled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) DatePickerMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) TimeZone() (*foundation.NSTimeZone, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDatePicker) TimeInterval() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

