//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDPowerStateRec
*/

type VDPowerStateRec struct {
  PowerState uint `v8:"powerState"`
  PowerFlags uint `v8:"powerFlags"`
  PowerReserved1 uint64 `v8:"powerReserved1"`
  PowerReserved2 uint64 `v8:"powerReserved2"`
}
