//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDDisplayConnectInfoRec
*/

type VDDisplayConnectInfoRec struct {
  CsDisplayType uint16 `v8:"csDisplayType"`
  CsConnectTaggedType byte `v8:"csConnectTaggedType"`
  CsConnectTaggedData byte `v8:"csConnectTaggedData"`
  CsConnectFlags uint `v8:"csConnectFlags"`
  CsDisplayComponent uint64 `v8:"csDisplayComponent"`
  CsConnectReserved uint64 `v8:"csConnectReserved"`
}
