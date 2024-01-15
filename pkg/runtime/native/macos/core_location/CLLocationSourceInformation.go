//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreLocation.CLLocationSourceInformation : objc.NSObject
*/

type CLLocationSourceInformation struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *CLLocationSourceInformation) InitWithSoftwareSimulationState() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLLocationSourceInformation) IsSimulatedBySoftware() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CLLocationSourceInformation) IsProducedByAccessory() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

