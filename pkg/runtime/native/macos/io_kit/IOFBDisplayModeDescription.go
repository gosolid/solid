//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOFBDisplayModeDescription
*/

type IOFBDisplayModeDescription struct {
  Info IODisplayModeInformation `v8:"info"`
  TimingInfo IOTimingInformation `v8:"timingInfo"`
}
