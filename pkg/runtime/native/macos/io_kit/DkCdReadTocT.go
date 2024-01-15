//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.dk_cd_read_toc_t
*/

type DkCdReadTocT struct {
  Format byte `v8:"format"`
  FormatAsTime byte `v8:"formatAsTime"`
  Reserved0016 [5]byte `v8:"reserved0016"`
  Address void `v8:"address"`
  Reserved0064 [6]byte `v8:"reserved0064"`
  BufferLength uint16 `v8:"bufferLength"`
  Buffer void `v8:"buffer"`
}
