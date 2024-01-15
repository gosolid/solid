//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/IOKit.framework/Headers/ndrvsupport/IOMacOSTypes.h:155:8)
*/

type Struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/IOKit.framework/Headers/ndrvsupport/IOMacOSTypes.h:155:8) struct {
  MajorRev byte `v8:"majorRev"`
  MinorAndBugRev byte `v8:"minorAndBugRev"`
  Stage byte `v8:"stage"`
  NonRelRev byte `v8:"nonRelRev"`
}
