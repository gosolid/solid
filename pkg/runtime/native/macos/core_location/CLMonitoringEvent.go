//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreLocation.CLMonitoringEvent : objc.NSObject
*/

type CLMonitoringEvent struct {
  *objc.NSObject
  *foundation.NSSecureCoding
}

func (r *CLMonitoringEvent) Refinement() (*CLCondition, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLMonitoringEvent) State() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CLMonitoringEvent) Date() (*foundation.NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLMonitoringEvent) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLMonitoringEvent) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLMonitoringEvent) Identifier() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

