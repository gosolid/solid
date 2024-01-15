//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
struct ApplicationServices.CMDeviceProfileInfo
*/

type CMDeviceProfileInfo struct {
  DataVersion uint `v8:"dataVersion"`
  ProfileID uint `v8:"profileID"`
  ProfileLoc CMProfileLocation `v8:"profileLoc"`
  ProfileName *core_foundation.CFDictionary `v8:"profileName"`
  Reserved uint `v8:"reserved"`
}
