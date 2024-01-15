//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSCalendar : objc.NSObject
*/

type NSCalendar struct {
  *objc.NSObject

}

func (r *NSCalendar) MinimumRangeOfUnit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) StartOfDayForDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) RangeOfWeekendStartDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) DateByAddingUnit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) CurrentCalendar() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) SetLocale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) GetHour() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) DateWithEra() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) NextWeekendStartDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) EnumerateDatesStartingAfterDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) Locale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) TimeZone() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) IsDateInWeekend() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) DateBySettingHour() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) LongEraSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) ShortMonthSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) ShortWeekdaySymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) StandaloneWeekdaySymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) ShortStandaloneWeekdaySymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) CalendarWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) DateFromComponents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) VeryShortMonthSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) StandaloneMonthSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) ShortQuarterSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) GetEra() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) IsDateInYesterday() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) SetMinimumDaysInFirstWeek() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) VeryShortWeekdaySymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) AMSymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) DateByAddingComponents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) SetTimeZone() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) Component() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) IsDateInToday() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) CalendarIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) OrdinalityOfUnit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) MaximumRangeOfUnit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) EraSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) VeryShortStandaloneMonthSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) PMSymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) RangeOfUnit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) IsDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) WeekdaySymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) NextDateAfterDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) DateBySettingUnit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) SetFirstWeekday() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) CompareDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) IsDateInTomorrow() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) Date() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) MonthSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) ShortStandaloneQuarterSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) MinimumDaysInFirstWeek() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) VeryShortStandaloneWeekdaySymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) StandaloneQuarterSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) ComponentsInTimeZone() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) ShortStandaloneMonthSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) InitWithCalendarIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) Components() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) AutoupdatingCurrentCalendar() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) FirstWeekday() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCalendar) QuarterSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

