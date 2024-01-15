//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.dk_cd_read_t
*/

type DkCdReadT struct {
  Offset uint64 `v8:"offset"`
  SectorArea byte `v8:"sectorArea"`
  SectorType byte `v8:"sectorType"`
  Reserved0080 [10]byte `v8:"reserved0080"`
  BufferLength uint `v8:"bufferLength"`
  Buffer void `v8:"buffer"`
}
