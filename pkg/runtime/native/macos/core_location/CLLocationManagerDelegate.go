//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreLocation.CLLocationManagerDelegate
*/

type CLLocationManagerDelegate struct {

}

func (r *CLLocationManagerDelegate) LocationManagerShouldDisplayHeadingCalibration() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CLLocationManagerDelegate) LocationManagerDidChangeAuthorization() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManagerDelegate) LocationManagerDidPauseLocationUpdates() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManagerDelegate) LocationManagerDidResumeLocationUpdates() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManagerDelegate) LocationManager() error {
  return fmt.Errorf("unimplemented")
}

