//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDGammaInfoRec
*/

type VDGammaInfoRec struct {
  CsLastGammaID uint `v8:"csLastGammaID"`
  CsNextGammaID uint `v8:"csNextGammaID"`
  CsGammaPtr *byte `v8:"csGammaPtr"`
  CsReserved uint64 `v8:"csReserved"`
}
