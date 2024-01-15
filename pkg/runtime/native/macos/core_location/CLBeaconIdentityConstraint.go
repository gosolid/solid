//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreLocation.CLBeaconIdentityConstraint : CoreLocation.CLBeaconIdentityCondition
*/

type CLBeaconIdentityConstraint struct {
  *CLBeaconIdentityCondition
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

