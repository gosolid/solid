//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSTimeZone : objc.NSObject
*/

type NSTimeZone struct {
  *objc.NSObject

}

func (r *NSTimeZone) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimeZone) Data() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimeZone) SecondsFromGMTForDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimeZone) AbbreviationForDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimeZone) IsDaylightSavingTimeForDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimeZone) DaylightSavingTimeOffsetForDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimeZone) NextDaylightSavingTimeTransitionAfterDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

