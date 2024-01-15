//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreLocation.CLBeaconIdentityCondition : CoreLocation.CLCondition
*/

type CLBeaconIdentityCondition struct {
  *CLCondition
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *CLBeaconIdentityCondition) Major() (*foundation.NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLBeaconIdentityCondition) Minor() (*foundation.NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLBeaconIdentityCondition) InitWithUUID() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLBeaconIdentityCondition) UUID() (*foundation.NSUUID, error) {
  return nil, fmt.Errorf("unimplemented")
}

