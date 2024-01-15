//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIDevice : objc.NSObject
*/

type UIDevice struct {
  *objc.NSObject

}

func (r *UIDevice) CurrentDevice() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDevice) SystemName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDevice) Orientation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDevice) IsBatteryMonitoringEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDevice) EndGeneratingDeviceOrientationNotifications() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDevice) PlayInputClick() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDevice) SetBatteryMonitoringEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDevice) SetProximityMonitoringEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDevice) UserInterfaceIdiom() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDevice) BeginGeneratingDeviceOrientationNotifications() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDevice) SystemVersion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDevice) IdentifierForVendor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDevice) IsGeneratingDeviceOrientationNotifications() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDevice) BatteryState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDevice) BatteryLevel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDevice) IsProximityMonitoringEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDevice) Model() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDevice) LocalizedModel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDevice) ProximityState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDevice) IsMultitaskingSupported() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDevice) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

