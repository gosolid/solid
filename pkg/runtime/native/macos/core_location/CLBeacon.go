//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreLocation.CLBeacon : objc.NSObject
*/

type CLBeacon struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *CLBeacon) Major() (*foundation.NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLBeacon) Minor() (*foundation.NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLBeacon) Proximity() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLBeacon) Accuracy() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLBeacon) Rssi() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLBeacon) Timestamp() (*foundation.NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLBeacon) UUID() (*foundation.NSUUID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLBeacon) ProximityUUID() (*foundation.NSUUID, error) {
  return nil, fmt.Errorf("unimplemented")
}

