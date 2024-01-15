//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/CoreServices.framework/Frameworks/CarbonCore.framework/Headers/fp.h:1320:5)
*/

type Struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/CoreServices.framework/Frameworks/CarbonCore.framework/Headers/fp.h:1320:5) struct {
  Length byte `v8:"length"`
  Text [36]byte `v8:"text"`
  Unused byte `v8:"unused"`
}
