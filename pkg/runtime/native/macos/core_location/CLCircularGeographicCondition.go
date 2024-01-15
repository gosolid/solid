//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreLocation.CLCircularGeographicCondition : CoreLocation.CLCondition
*/

type CLCircularGeographicCondition struct {
  *CLCondition
  *foundation.NSSecureCoding
}

func (r *CLCircularGeographicCondition) Radius() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLCircularGeographicCondition) InitWithCenter() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLCircularGeographicCondition) Center() (CLLocationCoordinate2D, error) {
  return CLLocationCoordinate2D{}, fmt.Errorf("unimplemented")
}

