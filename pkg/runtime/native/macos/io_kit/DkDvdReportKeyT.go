//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.dk_dvd_report_key_t
*/

type DkDvdReportKeyT struct {
  Format byte `v8:"format"`
  KeyClass byte `v8:"keyClass"`
  BlockCount byte `v8:"blockCount"`
  Reserved0024 [1]byte `v8:"reserved0024"`
  Address uint `v8:"address"`
  GrantID byte `v8:"grantID"`
  Reserved0072 [5]byte `v8:"reserved0072"`
  BufferLength uint16 `v8:"bufferLength"`
  Buffer void `v8:"buffer"`
}
