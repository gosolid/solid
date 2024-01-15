//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSDateFormatter : Foundation.NSFormatter
*/

type NSDateFormatter struct {
  *NSFormatter

}

func (r *NSDateFormatter) ShortMonthSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) DateFormatFromTemplate() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) FormattingContext() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetDateFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetVeryShortMonthSymbols() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) ShortStandaloneQuarterSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) StringFromDate() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetCalendar() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) VeryShortStandaloneMonthSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetVeryShortWeekdaySymbols() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetStandaloneQuarterSymbols() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) DateStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetLocale() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) WeekdaySymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetLongEraSymbols() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetDefaultFormatterBehavior() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) TimeZone() (*NSTimeZone, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) DefaultDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) LongEraSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) ShortStandaloneMonthSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) FormatterBehavior() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) TimeStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetMonthSymbols() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetShortMonthSymbols() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetVeryShortStandaloneWeekdaySymbols() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetFormattingContext() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetTimeZone() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetWeekdaySymbols() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) ShortQuarterSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetGregorianStartDate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetLocalizedDateFormatFromTemplate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) QuarterSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetAMSymbol() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetLenient() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetTwoDigitStartDate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) MonthSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) ShortStandaloneWeekdaySymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetShortStandaloneQuarterSymbols() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetDoesRelativeDateFormatting() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetDateStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) GeneratesCalendarDates() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) EraSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) ShortWeekdaySymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) StandaloneMonthSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetShortStandaloneWeekdaySymbols() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) DoesRelativeDateFormatting() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) DefaultFormatterBehavior() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetShortWeekdaySymbols() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) Locale() (*NSLocale, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) TwoDigitStartDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) PMSymbol() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) VeryShortWeekdaySymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) GregorianStartDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) LocalizedStringFromDate() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetTimeStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetFormatterBehavior() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) AMSymbol() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetShortStandaloneMonthSymbols() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetVeryShortStandaloneMonthSymbols() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetStandaloneWeekdaySymbols() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetQuarterSymbols() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) GetObjectValue() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetEraSymbols() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetPMSymbol() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetStandaloneMonthSymbols() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetShortQuarterSymbols() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) StandaloneQuarterSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) DateFromString() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) Calendar() (*NSCalendar, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) IsLenient() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) VeryShortMonthSymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) StandaloneWeekdaySymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) VeryShortStandaloneWeekdaySymbols() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetGeneratesCalendarDates() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetDefaultDate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) DateFormat() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

