//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSMeasurement : objc.NSObject
*/

type NSMeasurement struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSMeasurement) MeasurementByConvertingToUnit() (*NSMeasurement, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurement) MeasurementByAddingMeasurement() (*NSMeasurement, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurement) MeasurementBySubtractingMeasurement() (*NSMeasurement, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurement) Unit() (**NSUnit, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurement) DoubleValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMeasurement) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurement) InitWithDoubleValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurement) CanBeConvertedToUnit() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

