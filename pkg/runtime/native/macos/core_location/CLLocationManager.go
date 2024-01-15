//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreLocation.CLLocationManager : objc.NSObject
*/

type CLLocationManager struct {
  *objc.NSObject

}

func (r *CLLocationManager) StopMonitoringLocationPushes() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) StopMonitoringForRegion() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) StartRangingBeaconsInRegion() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) ShowsBackgroundLocationIndicator() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) Heading() (*CLHeading, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) StopUpdatingLocation() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) SetActivityType() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) DistanceFilter() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) DeferredLocationUpdatesAvailable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) StopUpdatingHeading() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) SetPausesLocationUpdatesAutomatically() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) StartUpdatingHeading() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) RequestLocation() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) AllowDeferredLocationUpdatesUntilTraveled() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) DisallowDeferredLocationUpdates() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) ActivityType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) AllowsBackgroundLocationUpdates() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) AuthorizationStatus() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) HeadingOrientation() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) SetDistanceFilter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) StartRangingBeaconsSatisfyingConstraint() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) RangedRegions() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) RangedBeaconConstraints() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) StopMonitoringSignificantLocationChanges() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) Location() (*CLLocation, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) HeadingFilter() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) SetHeadingOrientation() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) StartUpdatingLocation() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) DismissHeadingCalibrationDisplay() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) DesiredAccuracy() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) SetDesiredAccuracy() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) RequestWhenInUseAuthorization() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) StartMonitoringLocationPushesWithCompletion() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) RequestAlwaysAuthorization() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) RequestStateForRegion() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) AccuracyAuthorization() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) SetPurpose() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) MaximumRegionMonitoringDistance() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) StartMonitoringForRegion() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) IsMonitoringAvailableForClass() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) StopRangingBeaconsSatisfyingConstraint() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) HeadingAvailable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) RegionMonitoringEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) IsRangingAvailable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) Purpose() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) SetShowsBackgroundLocationIndicator() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) SignificantLocationChangeMonitoringAvailable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) StartMonitoringSignificantLocationChanges() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) StopRangingBeaconsInRegion() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) SetHeadingFilter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) RegionMonitoringAvailable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) RequestTemporaryFullAccuracyAuthorizationWithPurposeKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) RequestHistoricalLocationsWithPurposeKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) IsAuthorizedForWidgetUpdates() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) LocationServicesEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) SetAllowsBackgroundLocationUpdates() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) PausesLocationUpdatesAutomatically() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CLLocationManager) MonitoredRegions() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

