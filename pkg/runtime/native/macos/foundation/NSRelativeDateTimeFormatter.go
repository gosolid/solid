//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSRelativeDateTimeFormatter : Foundation.NSFormatter
*/

type NSRelativeDateTimeFormatter struct {
  *NSFormatter

}

func (r *NSRelativeDateTimeFormatter) DateTimeStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) Calendar() (*NSCalendar, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) LocalizedStringForDate() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) StringForObjectValue() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) UnitsStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) FormattingContext() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) LocalizedStringFromTimeInterval() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) SetDateTimeStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) SetUnitsStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) SetFormattingContext() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) SetCalendar() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) Locale() (*NSLocale, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) SetLocale() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) LocalizedStringFromDateComponents() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

