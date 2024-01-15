//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
struct ApplicationServices.CMDeviceProfileScope
*/

type CMDeviceProfileScope struct {
  DeviceUser *core_foundation.CFString `v8:"deviceUser"`
  DeviceHost *core_foundation.CFString `v8:"deviceHost"`
}
