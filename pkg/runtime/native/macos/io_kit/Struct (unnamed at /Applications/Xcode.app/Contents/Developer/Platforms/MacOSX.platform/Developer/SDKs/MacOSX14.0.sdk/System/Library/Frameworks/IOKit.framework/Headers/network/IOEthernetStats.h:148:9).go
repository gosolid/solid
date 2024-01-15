//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/IOKit.framework/Headers/network/IOEthernetStats.h:148:9)
*/

type Struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/IOKit.framework/Headers/network/IOEthernetStats.h:148:9) struct {
  Dot3StatsEntry Struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/IOKit.framework/Headers/network/IOEthernetStats.h:55:9) `v8:"dot3StatsEntry"`
  Dot3CollEntry Struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/IOKit.framework/Headers/network/IOEthernetStats.h:79:9) `v8:"dot3CollEntry"`
  Dot3RxExtraEntry Struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/IOKit.framework/Headers/network/IOEthernetStats.h:99:9) `v8:"dot3RxExtraEntry"`
  Dot3TxExtraEntry Struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/IOKit.framework/Headers/network/IOEthernetStats.h:126:9) `v8:"dot3TxExtraEntry"`
}
