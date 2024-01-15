//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.dk_cd_read_isrc_t
*/

type DkCdReadIsrcT struct {
  Isrc [13]byte `v8:"isrc"`
  Track byte `v8:"track"`
  Reserved0112 [2]byte `v8:"reserved0112"`
}
