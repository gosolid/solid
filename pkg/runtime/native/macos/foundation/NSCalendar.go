//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSCalendar : objc.NSObject
*/

type NSCalendar struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSCalendar) InitWithCalendarIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) DateBySettingHour() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) ShortMonthSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) StandaloneMonthSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) MinimumRangeOfUnit() (NSRange, error) {
  return NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) Component() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) Locale() (*NSLocale, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) FirstWeekday() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) EraSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) VeryShortStandaloneWeekdaySymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) ShortStandaloneMonthSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) ShortWeekdaySymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) ShortStandaloneWeekdaySymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) GetEra() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCalendar) RangeOfWeekendStartDate() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) DateByAddingUnit() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) CurrentCalendar() (*NSCalendar, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) ShortQuarterSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) PMSymbol() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) OrdinalityOfUnit() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) IsDateInWeekend() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) DateBySettingUnit() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) MinimumDaysInFirstWeek() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) StandaloneWeekdaySymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) IsDateInTomorrow() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) DateFromComponents() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) DateWithEra() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) TimeZone() (*NSTimeZone, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) CalendarWithIdentifier() (*NSCalendar, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) IsDateInToday() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) SetLocale() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCalendar) QuarterSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) CalendarIdentifier() (**NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) SetMinimumDaysInFirstWeek() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCalendar) VeryShortWeekdaySymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) SetFirstWeekday() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCalendar) VeryShortStandaloneMonthSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) IsDate() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) NextDateAfterDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) AutoupdatingCurrentCalendar() (*NSCalendar, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) StandaloneQuarterSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) DateByAddingComponents() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) StartOfDayForDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) NextWeekendStartDate() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) MaximumRangeOfUnit() (NSRange, error) {
  return NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) VeryShortMonthSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) ShortStandaloneQuarterSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) AMSymbol() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) SetTimeZone() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCalendar) MonthSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) WeekdaySymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) ComponentsInTimeZone() (*NSDateComponents, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) IsDateInYesterday() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) LongEraSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) RangeOfUnit() (NSRange, error) {
  return NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) Components() (*NSDateComponents, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) GetHour() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCalendar) CompareDate() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) EnumerateDatesStartingAfterDate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCalendar) Date() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

