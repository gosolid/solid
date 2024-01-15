//js:package native/macos/core-location
package core_location

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreLocation.CLMonitor : objc.NSObject
*/

type CLMonitor struct {
  *objc.NSObject

}

func (r *CLMonitor) RequestMonitorWithConfiguration() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLMonitor) AddConditionForMonitoring() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLMonitor) RemoveConditionFromMonitoringWithIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *CLMonitor) MonitoringRecordForIdentifier() (*CLMonitoringRecord, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLMonitor) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLMonitor) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLMonitor) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CLMonitor) MonitoredIdentifiers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

