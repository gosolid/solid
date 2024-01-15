//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreLocation.CLMonitoringRecord : objc.NSObject
*/

type CLMonitoringRecord struct {
  *objc.NSObject
  *foundation.NSSecureCoding
}

func (r *CLMonitoringRecord) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLMonitoringRecord) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLMonitoringRecord) Condition() (*CLCondition, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLMonitoringRecord) LastEvent() (*CLMonitoringEvent, error) {
  return nil, fmt.Errorf("unimplemented")
}

