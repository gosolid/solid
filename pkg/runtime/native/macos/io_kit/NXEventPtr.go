//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.NXEventPtr
*/

type NXEventPtr struct {
  Type int `v8:"type"`
  Location Struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/IOKit.framework/Headers/hidsystem/IOLLEvent.h:505:5) `v8:"location"`
  Time uint64 `v8:"time"`
  Flags int `v8:"flags"`
  Window uint `v8:"window"`
  ServiceId uint64 `v8:"serviceId"`
  ExtPid int `v8:"extPid"`
  Data void `v8:"data"`
}
