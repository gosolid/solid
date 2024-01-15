//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDGetGammaListRec
*/

type VDGetGammaListRec struct {
  CsPreviousGammaTableID uint `v8:"csPreviousGammaTableID"`
  CsGammaTableID uint `v8:"csGammaTableID"`
  CsGammaTableSize uint `v8:"csGammaTableSize"`
  CsGammaTableName byte `v8:"csGammaTableName"`
}
