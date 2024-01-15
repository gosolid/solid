//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSDateIntervalFormatter : Foundation.NSFormatter
*/

type NSDateIntervalFormatter struct {
  *NSFormatter

}

func (r *NSDateIntervalFormatter) Calendar() (*NSCalendar, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) TimeZone() (*NSTimeZone, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) DateStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) SetTimeStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) StringFromDateInterval() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) SetTimeZone() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) Locale() (*NSLocale, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) DateTemplate() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) SetDateTemplate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) SetDateStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) StringFromDate() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) SetLocale() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) SetCalendar() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) TimeStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

