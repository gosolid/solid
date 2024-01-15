//js:package native/ios/foundation
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

}

func (r *NSMeasurementFormatter) StringFromUnit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurementFormatter) UnitStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurementFormatter) SetUnitStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurementFormatter) SetLocale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurementFormatter) NumberFormatter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurementFormatter) StringFromMeasurement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurementFormatter) SetUnitOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurementFormatter) Locale() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurementFormatter) SetNumberFormatter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurementFormatter) UnitOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

