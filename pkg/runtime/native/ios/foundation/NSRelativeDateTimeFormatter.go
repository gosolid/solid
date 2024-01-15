//js:package native/ios/foundation
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

func (r *NSRelativeDateTimeFormatter) StringForObjectValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) SetUnitsStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) FormattingContext() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) Locale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) LocalizedStringFromTimeInterval() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) DateTimeStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) SetFormattingContext() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) SetCalendar() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) Calendar() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) SetLocale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) LocalizedStringFromDateComponents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) LocalizedStringForDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) SetDateTimeStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRelativeDateTimeFormatter) UnitsStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

