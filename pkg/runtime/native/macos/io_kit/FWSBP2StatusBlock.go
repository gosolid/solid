//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.FWSBP2StatusBlock
*/

type FWSBP2StatusBlock struct {
  Details byte `v8:"details"`
  SbpStatus byte `v8:"sbpStatus"`
  OrbOffsetHi uint16 `v8:"orbOffsetHi"`
  OrbOffsetLo uint `v8:"orbOffsetLo"`
  Status [6]uint `v8:"status"`
}
