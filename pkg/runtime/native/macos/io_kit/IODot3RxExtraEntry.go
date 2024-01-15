//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IODot3RxExtraEntry
*/

type IODot3RxExtraEntry struct {
  Overruns uint `v8:"overruns"`
  WatchdogTimeouts uint `v8:"watchdogTimeouts"`
  FrameTooShorts uint `v8:"frameTooShorts"`
  CollisionErrors uint `v8:"collisionErrors"`
  PhyErrors uint `v8:"phyErrors"`
  Timeouts uint `v8:"timeouts"`
  Interrupts uint `v8:"interrupts"`
  Resets uint `v8:"resets"`
  ResourceErrors uint `v8:"resourceErrors"`
  Reserved [4]uint `v8:"reserved"`
}
