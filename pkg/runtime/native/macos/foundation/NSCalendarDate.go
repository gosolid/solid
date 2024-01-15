//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSCalendarDate : Foundation.NSDate
*/

type NSCalendarDate struct {
  *NSDate

}

func (r *NSCalendarDate) DayOfMonth() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCalendarDate) DayOfWeek() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCalendarDate) YearOfCommonEra() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCalendarDate) TimeZone() (*NSTimeZone, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendarDate) InitWithYear() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendarDate) SetTimeZone() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCalendarDate) DistantPast() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendarDate) DateWithString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendarDate) DateWithYear() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendarDate) DayOfYear() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCalendarDate) MinuteOfHour() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCalendarDate) MonthOfYear() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCalendarDate) DescriptionWithCalendarFormat() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendarDate) InitWithString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendarDate) CalendarDate() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendarDate) DayOfCommonEra() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCalendarDate) SecondOfMinute() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCalendarDate) CalendarFormat() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendarDate) DescriptionWithLocale() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendarDate) Years() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCalendarDate) DateByAddingYears() (*NSCalendarDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendarDate) HourOfDay() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCalendarDate) SetCalendarFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCalendarDate) DistantFuture() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

