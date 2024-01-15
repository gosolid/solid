//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDMirrorRec
*/

type VDMirrorRec struct {
  CsMirrorSize uint `v8:"csMirrorSize"`
  CsMirrorVersion uint `v8:"csMirrorVersion"`
  CsMirrorRequestID RegEntryID `v8:"csMirrorRequestID"`
  CsMirrorResultID RegEntryID `v8:"csMirrorResultID"`
  CsMirrorFeatures uint `v8:"csMirrorFeatures"`
  CsMirrorSupportedFlags uint `v8:"csMirrorSupportedFlags"`
  CsMirrorFlags uint `v8:"csMirrorFlags"`
  CsReserved1 uint `v8:"csReserved1"`
  CsReserved2 uint `v8:"csReserved2"`
  CsReserved3 uint `v8:"csReserved3"`
  CsReserved4 uint `v8:"csReserved4"`
  CsReserved5 uint `v8:"csReserved5"`
}
