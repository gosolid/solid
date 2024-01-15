//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDSetHardwareCursorRec
*/

type VDSetHardwareCursorRec struct {
  CsCursorRef void `v8:"csCursorRef"`
  CsReserved1 uint `v8:"csReserved1"`
  CsReserved2 uint `v8:"csReserved2"`
}
