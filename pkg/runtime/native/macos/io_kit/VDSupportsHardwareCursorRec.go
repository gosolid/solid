//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDSupportsHardwareCursorRec
*/

type VDSupportsHardwareCursorRec struct {
  CsSupportsHardwareCursor uint `v8:"csSupportsHardwareCursor"`
  CsReserved1 uint `v8:"csReserved1"`
  CsReserved2 uint `v8:"csReserved2"`
}
