//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreLocation.CLHeading : objc.NSObject
*/

type CLHeading struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *CLHeading) Y() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLHeading) Z() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLHeading) Timestamp() (*foundation.NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLHeading) MagneticHeading() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLHeading) TrueHeading() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLHeading) HeadingAccuracy() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLHeading) X() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

