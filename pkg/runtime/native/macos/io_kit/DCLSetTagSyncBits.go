//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.DCLSetTagSyncBits
*/

type DCLSetTagSyncBits struct {
  PNextDCLCommand DCLCommand `v8:"pNextDCLCommand"`
  CompilerData uint `v8:"compilerData"`
  Opcode uint `v8:"opcode"`
  TagBits uint16 `v8:"tagBits"`
  SyncBits uint16 `v8:"syncBits"`
}
