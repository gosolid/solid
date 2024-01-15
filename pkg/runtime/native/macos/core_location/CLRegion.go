//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface CoreLocation.CLRegion : objc.NSObject
*/

type CLRegion struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *CLRegion) SetNotifyOnExit() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLRegion) InitCircularRegionWithCenter() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLRegion) ContainsCoordinate() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CLRegion) Center() (CLLocationCoordinate2D, error) {
  return CLLocationCoordinate2D{}, fmt.Errorf("unimplemented")
}

func (r *CLRegion) SetNotifyOnEntry() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLRegion) NotifyOnExit() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CLRegion) Radius() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLRegion) Identifier() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLRegion) NotifyOnEntry() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

