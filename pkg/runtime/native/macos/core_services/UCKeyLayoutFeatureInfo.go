//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.UCKeyLayoutFeatureInfo
*/

type UCKeyLayoutFeatureInfo struct {
  KeyLayoutFeatureInfoFormat uint16 `v8:"keyLayoutFeatureInfoFormat"`
  Reserved uint16 `v8:"reserved"`
  MaxOutputStringLength uint `v8:"maxOutputStringLength"`
}
