//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDDDCBlockRec
*/

type VDDDCBlockRec struct {
  DdcBlockNumber uint `v8:"ddcBlockNumber"`
  DdcBlockType uint `v8:"ddcBlockType"`
  DdcFlags uint `v8:"ddcFlags"`
  DdcReserved uint `v8:"ddcReserved"`
  DdcBlockData [128]byte `v8:"ddcBlockData"`
}
