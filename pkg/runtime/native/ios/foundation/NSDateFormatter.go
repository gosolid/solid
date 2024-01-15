//js:package native/ios/foundation
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

func (r *NSDateFormatter) TimeStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) FormatterBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetLenient() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) DefaultDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetShortMonthSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetStandaloneQuarterSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) DateStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetTimeStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) Calendar() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) EraSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetPMSymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) Locale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) ShortMonthSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) WeekdaySymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) ShortStandaloneWeekdaySymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) GregorianStartDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetDoesRelativeDateFormatting() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) TwoDigitStartDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetAMSymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetStandaloneMonthSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) StandaloneQuarterSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) StandaloneWeekdaySymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetFormattingContext() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetLocale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) IsLenient() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) MonthSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetWeekdaySymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetLongEraSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) VeryShortStandaloneMonthSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetLocalizedDateFormatFromTemplate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetDateStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetTwoDigitStartDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetShortWeekdaySymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) StandaloneMonthSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetVeryShortStandaloneWeekdaySymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) StringFromDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetTimeZone() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) DoesRelativeDateFormatting() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) GetObjectValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) ShortStandaloneMonthSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetStandaloneWeekdaySymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetQuarterSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) ShortQuarterSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) FormattingContext() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) AMSymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetShortStandaloneMonthSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetVeryShortStandaloneMonthSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) LongEraSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) VeryShortMonthSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) DefaultFormatterBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetDateFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetFormatterBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) QuarterSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) DateFromString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) DateFormatFromTemplate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetGeneratesCalendarDates() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetEraSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) VeryShortWeekdaySymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetShortStandaloneWeekdaySymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetDefaultFormatterBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetDefaultDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetVeryShortWeekdaySymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) ShortStandaloneQuarterSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetShortStandaloneQuarterSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) DateFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) GeneratesCalendarDates() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetCalendar() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) ShortWeekdaySymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) PMSymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetVeryShortMonthSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetShortQuarterSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetGregorianStartDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) LocalizedStringFromDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) TimeZone() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) SetMonthSymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDateFormatter) VeryShortStandaloneWeekdaySymbols() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

