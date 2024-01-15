//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IODot3TxExtraEntry
*/

type IODot3TxExtraEntry struct {
  Underruns uint `v8:"underruns"`
  Jabbers uint `v8:"jabbers"`
  PhyErrors uint `v8:"phyErrors"`
  Timeouts uint `v8:"timeouts"`
  Interrupts uint `v8:"interrupts"`
  Resets uint `v8:"resets"`
  ResourceErrors uint `v8:"resourceErrors"`
  Reserved [4]uint `v8:"reserved"`
}
