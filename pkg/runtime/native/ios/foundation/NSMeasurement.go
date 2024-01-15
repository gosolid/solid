//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSMeasurement : objc.NSObject
*/

type NSMeasurement struct {
  *objc.NSObject

}

func (r *NSMeasurement) Unit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurement) DoubleValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurement) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurement) InitWithDoubleValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurement) CanBeConvertedToUnit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurement) MeasurementByConvertingToUnit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurement) MeasurementByAddingMeasurement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMeasurement) MeasurementBySubtractingMeasurement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

