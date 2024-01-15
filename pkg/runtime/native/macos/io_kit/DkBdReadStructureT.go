//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.dk_bd_read_structure_t
*/

type DkBdReadStructureT struct {
  Format byte `v8:"format"`
  Reserved0008 [3]byte `v8:"reserved0008"`
  Address uint `v8:"address"`
  GrantID byte `v8:"grantID"`
  Layer byte `v8:"layer"`
  Reserved0080 [4]byte `v8:"reserved0080"`
  BufferLength uint16 `v8:"bufferLength"`
  Buffer void `v8:"buffer"`
}
