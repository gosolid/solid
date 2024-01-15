//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/IOKit.framework/Headers/ndrvsupport/IOMacOSTypes.h:125:8)
*/

type Struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/IOKit.framework/Headers/ndrvsupport/IOMacOSTypes.h:125:8) struct {
  Top int16 `v8:"top"`
  Left int16 `v8:"left"`
  Bottom int16 `v8:"bottom"`
  Right int16 `v8:"right"`
}
