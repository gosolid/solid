//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.DriverOSRuntime
*/

type DriverOSRuntime struct {
  DriverRuntime uint `v8:"driverRuntime"`
  DriverName [32]byte `v8:"driverName"`
  DriverDescReserved [8]uint `v8:"driverDescReserved"`
}
