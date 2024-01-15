//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreLocation.CLBeaconRegion : CoreLocation.CLRegion
*/

type CLBeaconRegion struct {
  *CLRegion

}

func (r *CLBeaconRegion) InitWithBeaconIdentityConstraint() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLBeaconRegion) PeripheralDataWithMeasuredPower() (*foundation.NSMutableDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLBeaconRegion) Major() (*foundation.NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLBeaconRegion) SetNotifyEntryStateOnDisplay() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLBeaconRegion) InitWithUUID() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLBeaconRegion) InitWithProximityUUID() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLBeaconRegion) BeaconIdentityConstraint() (*CLBeaconIdentityConstraint, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLBeaconRegion) UUID() (*foundation.NSUUID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLBeaconRegion) ProximityUUID() (*foundation.NSUUID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLBeaconRegion) Minor() (*foundation.NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLBeaconRegion) NotifyEntryStateOnDisplay() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

