//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreLocation.CLLocation : objc.NSObject
*/

type CLLocation struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *CLLocation) DistanceFromLocation() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLLocation) Coordinate() (CLLocationCoordinate2D, error) {
  return CLLocationCoordinate2D{}, fmt.Errorf("unimplemented")
}

func (r *CLLocation) EllipsoidalAltitude() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLLocation) HorizontalAccuracy() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLLocation) Course() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLLocation) Speed() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLLocation) GetDistanceFrom() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLLocation) Altitude() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLLocation) VerticalAccuracy() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLLocation) CourseAccuracy() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLLocation) SpeedAccuracy() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLLocation) InitWithLatitude() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLLocation) Timestamp() (*foundation.NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLLocation) Floor() (*CLFloor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLLocation) SourceInformation() (*CLLocationSourceInformation, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLLocation) InitWithCoordinate() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

