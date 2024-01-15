//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.NXTabletPointDataPtr
*/

type NXTabletPointDataPtr struct {
  X int `v8:"x"`
  Y int `v8:"y"`
  Z int `v8:"z"`
  Buttons uint16 `v8:"buttons"`
  Pressure uint16 `v8:"pressure"`
  Tilt Struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/IOKit.framework/Headers/hidsystem/IOLLEvent.h:351:5) `v8:"tilt"`
  Rotation uint16 `v8:"rotation"`
  TangentialPressure int16 `v8:"tangentialPressure"`
  DeviceID uint16 `v8:"deviceID"`
  Vendor1 int16 `v8:"vendor1"`
  Vendor2 int16 `v8:"vendor2"`
  Vendor3 int16 `v8:"vendor3"`
}
