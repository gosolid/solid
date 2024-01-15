//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSTimeZone : objc.NSObject
*/

type NSTimeZone struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSTimeZone) AbbreviationForDate() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimeZone) IsDaylightSavingTimeForDate() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTimeZone) DaylightSavingTimeOffsetForDate() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTimeZone) NextDaylightSavingTimeTransitionAfterDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimeZone) Name() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimeZone) Data() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimeZone) SecondsFromGMTForDate() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

