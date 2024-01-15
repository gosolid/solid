//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CoreLocation.CLCircularRegion : CoreLocation.CLRegion
*/

type CLCircularRegion struct {
  *CLRegion

}

func (r *CLCircularRegion) InitWithCenter() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLCircularRegion) ContainsCoordinate() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CLCircularRegion) Center() (CLLocationCoordinate2D, error) {
  return CLLocationCoordinate2D{}, fmt.Errorf("unimplemented")
}

func (r *CLCircularRegion) Radius() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

