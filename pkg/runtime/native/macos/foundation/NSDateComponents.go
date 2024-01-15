//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSDateComponents : objc.NSObject
*/

type NSDateComponents struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSDateComponents) SetWeekOfMonth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) IsLeapMonth() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) Date() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) Era() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) Hour() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) Minute() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) WeekOfYear() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) SetYearForWeekOfYear() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) SetLeapMonth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) SetWeek() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) TimeZone() (*NSTimeZone, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) SetYear() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) SetSecond() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) Nanosecond() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) SetValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) Month() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) Second() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) Day() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) SetDay() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) Quarter() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) Calendar() (*NSCalendar, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) SetCalendar() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) SetTimeZone() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) SetQuarter() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) Week() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) SetMinute() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) SetNanosecond() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) IsValidDateInCalendar() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) SetWeekdayOrdinal() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) WeekdayOrdinal() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) SetWeekOfYear() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) YearForWeekOfYear() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) SetMonth() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) SetHour() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) Weekday() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) SetWeekday() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) WeekOfMonth() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) ValueForComponent() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) SetEra() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) Year() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateComponents) IsValidDate() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

