//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDHardwareCursorDrawStateRec
*/

type VDHardwareCursorDrawStateRec struct {
  CsCursorX int `v8:"csCursorX"`
  CsCursorY int `v8:"csCursorY"`
  CsCursorVisible uint `v8:"csCursorVisible"`
  CsCursorSet uint `v8:"csCursorSet"`
  CsReserved1 uint `v8:"csReserved1"`
  CsReserved2 uint `v8:"csReserved2"`
}
