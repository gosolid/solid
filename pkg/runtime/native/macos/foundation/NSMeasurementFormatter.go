//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSMeasurementFormatter : Foundation.NSFormatter
*/

type NSMeasurementFormatter struct {
  *NSFormatter
  *NSSecureCoding
}

func (r *NSMeasurementFormatter) SetUnitStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMeasurementFormatter) SetLocale() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMeasurementFormatter) SetNumberFormatter() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMeasurementFormatter) UnitOptions() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMeasurementFormatter) SetUnitOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMeasurementFormatter) UnitStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMeasurementFormatter) NumberFormatter() (*NSNumberFormatter, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurementFormatter) StringFromMeasurement() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurementFormatter) StringFromUnit() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurementFormatter) Locale() (*NSLocale, error) {
  return nil, fmt.Errorf("unimplemented")
}

