//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOPMSystemCapabilityChangeParameters
*/

type IOPMSystemCapabilityChangeParameters struct {
  NotifyRef uint `v8:"notifyRef"`
  MaxWaitForReply uint `v8:"maxWaitForReply"`
  ChangeFlags uint `v8:"changeFlags"`
  Reserved1 uint `v8:"reserved1"`
  FromCapabilities uint `v8:"fromCapabilities"`
  ToCapabilities uint `v8:"toCapabilities"`
  Reserved2 [4]uint `v8:"reserved2"`
}
