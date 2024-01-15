//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.DPME
*/

type DPME struct {
  DpmeSignature uint16 `v8:"dpmeSignature"`
  DpmeReserved1 uint16 `v8:"dpmeReserved1"`
  DpmeMapEntries uint `v8:"dpmeMapEntries"`
  DpmePblockStart uint `v8:"dpmePblockStart"`
  DpmePblocks uint `v8:"dpmePblocks"`
  DpmeName [32]byte `v8:"dpmeName"`
  DpmeType [32]byte `v8:"dpmeType"`
  DpmeLblockStart uint `v8:"dpmeLblockStart"`
  DpmeLblocks uint `v8:"dpmeLblocks"`
  DpmeFlags uint `v8:"dpmeFlags"`
  DpmeBootBlock uint `v8:"dpmeBootBlock"`
  DpmeBootBytes uint `v8:"dpmeBootBytes"`
  DpmeLoadAddr uint `v8:"dpmeLoadAddr"`
  DpmeLoadAddr2 uint `v8:"dpmeLoadAddr2"`
  DpmeGotoAddr uint `v8:"dpmeGotoAddr"`
  DpmeGotoAddr2 uint `v8:"dpmeGotoAddr2"`
  DpmeChecksum uint `v8:"dpmeChecksum"`
  DpmeProcessId [16]byte `v8:"dpmeProcessId"`
  DpmeReserved2 [32]uint `v8:"dpmeReserved2"`
  DpmeReserved3 [62]uint `v8:"dpmeReserved3"`
}
