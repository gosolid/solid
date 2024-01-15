//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/IOKit.framework/Headers/network/IOEthernetStats.h:126:9)
*/

type Struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/IOKit.framework/Headers/network/IOEthernetStats.h:126:9) struct {
  Underruns int `v8:"underruns"`
  Jabbers int `v8:"jabbers"`
  PhyErrors int `v8:"phyErrors"`
  Timeouts int `v8:"timeouts"`
  Interrupts int `v8:"interrupts"`
  Resets int `v8:"resets"`
  ResourceErrors int `v8:"resourceErrors"`
  Reserved int `v8:"reserved"`
}
