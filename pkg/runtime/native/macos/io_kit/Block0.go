//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.Block0
*/

type Block0 struct {
  SbSig uint16 `v8:"sbSig"`
  SbBlkSize uint16 `v8:"sbBlkSize"`
  SbBlkCount uint `v8:"sbBlkCount"`
  SbDevType uint16 `v8:"sbDevType"`
  SbDevId uint16 `v8:"sbDevId"`
  SbDrvrData uint `v8:"sbDrvrData"`
  SbDrvrCount uint16 `v8:"sbDrvrCount"`
  SbDrvrMap [8]DDMap `v8:"sbDrvrMap"`
  SbReserved [430]byte `v8:"sbReserved"`
}
