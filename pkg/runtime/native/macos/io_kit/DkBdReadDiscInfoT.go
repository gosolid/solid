//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.dk_bd_read_disc_info_t
*/

type DkBdReadDiscInfoT struct {
  Reserved0000 [14]byte `v8:"reserved0000"`
  BufferLength uint16 `v8:"bufferLength"`
  Buffer void `v8:"buffer"`
}
