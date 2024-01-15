//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.CDAudioStatus
*/

type CDAudioStatus struct {
  Status byte `v8:"status"`
  Position Struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/IOKit.framework/Headers/storage/IOCDTypes.h:76:5) `v8:"position"`
}
