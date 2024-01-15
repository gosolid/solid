//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.dk_dvd_read_rzone_info_t
*/

type DkDvdReadRzoneInfoT struct {
  Reserved0000 [4]byte `v8:"reserved0000"`
  Address uint `v8:"address"`
  AddressType byte `v8:"addressType"`
  Reserved0072 [5]byte `v8:"reserved0072"`
  BufferLength uint16 `v8:"bufferLength"`
  Buffer void `v8:"buffer"`
}
