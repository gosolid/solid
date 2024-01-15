//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSDatePickerCell : AppKit.NSActionCell
*/

type NSDatePickerCell struct {
  *NSActionCell

}

func (r *NSDatePickerCell) SetLocale() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) TimeInterval() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) SetMinDate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) DrawsBackground() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) SetDatePickerMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) DatePickerElements() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) SetTimeInterval() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) BackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) SetCalendar() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) SetDateValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) SetDrawsBackground() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) SetTextColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) DateValue() (*foundation.NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) SetBackgroundColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) Calendar() (*foundation.NSCalendar, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) DatePickerMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) SetDatePickerElements() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) MinDate() (*foundation.NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) SetMaxDate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) InitImageCell() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) DatePickerStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) SetDatePickerStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) TextColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) Locale() (*foundation.NSLocale, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) SetTimeZone() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) InitTextCell() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) TimeZone() (*foundation.NSTimeZone, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDatePickerCell) MaxDate() (*foundation.NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

