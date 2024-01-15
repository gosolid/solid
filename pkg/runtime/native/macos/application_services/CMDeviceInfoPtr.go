//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
struct ApplicationServices.CMDeviceInfoPtr
*/

type CMDeviceInfoPtr struct {
  DataVersion uint `v8:"dataVersion"`
  DeviceClass uint `v8:"deviceClass"`
  DeviceID uint `v8:"deviceID"`
  DeviceScope CMDeviceProfileScope `v8:"deviceScope"`
  DeviceState uint `v8:"deviceState"`
  DefaultProfileID uint `v8:"defaultProfileID"`
  DeviceName *core_foundation.CFDictionary `v8:"deviceName"`
  ProfileCount uint `v8:"profileCount"`
  Reserved uint `v8:"reserved"`
}
