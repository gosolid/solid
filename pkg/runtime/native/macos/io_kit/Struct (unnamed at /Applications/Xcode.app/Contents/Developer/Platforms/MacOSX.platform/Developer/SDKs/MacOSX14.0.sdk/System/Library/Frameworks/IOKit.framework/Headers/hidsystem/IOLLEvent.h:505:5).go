//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/IOKit.framework/Headers/hidsystem/IOLLEvent.h:505:5)
*/

type Struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/IOKit.framework/Headers/hidsystem/IOLLEvent.h:505:5) struct {
  X int `v8:"x"`
  Y int `v8:"y"`
}
