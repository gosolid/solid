//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.dk_cd_read_mcn_t
*/

type DkCdReadMcnT struct {
  Mcn [14]byte `v8:"mcn"`
  Reserved0112 [2]byte `v8:"reserved0112"`
}
