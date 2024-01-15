//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.dk_cd_read_disc_info_t
*/

type DkCdReadDiscInfoT struct {
  Reserved0000 [14]byte `v8:"reserved0000"`
  BufferLength uint16 `v8:"bufferLength"`
  Buffer void `v8:"buffer"`
}
