//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.dk_bd_send_key_t
*/

type DkBdSendKeyT struct {
  Format byte `v8:"format"`
  KeyClass byte `v8:"keyClass"`
  Reserved0016 [6]byte `v8:"reserved0016"`
  GrantID byte `v8:"grantID"`
  Reserved0072 [5]byte `v8:"reserved0072"`
  BufferLength uint16 `v8:"bufferLength"`
  Buffer void `v8:"buffer"`
}
