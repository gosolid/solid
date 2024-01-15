//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.DriverDescription
*/

type DriverDescription struct {
  DriverDescSignature uint `v8:"driverDescSignature"`
  DriverDescVersion uint `v8:"driverDescVersion"`
  DriverType DriverType `v8:"driverType"`
  DriverOSRuntimeInfo DriverOSRuntime `v8:"driverOSRuntimeInfo"`
  DriverServices DriverOSService `v8:"driverServices"`
}
