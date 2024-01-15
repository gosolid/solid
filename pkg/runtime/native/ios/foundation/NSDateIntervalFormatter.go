//js:package native/ios/foundation
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

func (r *NSDateIntervalFormatter) StringFromDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) SetCalendar() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) SetDateStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) SetLocale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) Calendar() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) SetTimeZone() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) SetTimeStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) DateTemplate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) DateStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) StringFromDateInterval() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) Locale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) TimeZone() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) SetDateTemplate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateIntervalFormatter) TimeStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

