//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreLocation.CLVisit : objc.NSObject
*/

type CLVisit struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *CLVisit) ArrivalDate() (*foundation.NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLVisit) DepartureDate() (*foundation.NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLVisit) Coordinate() (CLLocationCoordinate2D, error) {
  return CLLocationCoordinate2D{}, fmt.Errorf("unimplemented")
}

func (r *CLVisit) HorizontalAccuracy() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

