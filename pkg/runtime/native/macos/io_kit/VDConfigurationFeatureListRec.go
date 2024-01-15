//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDConfigurationFeatureListRec
*/

type VDConfigurationFeatureListRec struct {
  CsConfigFeatureList uint `v8:"csConfigFeatureList"`
  CsNumConfigFeatures uint64 `v8:"csNumConfigFeatures"`
  CsReserved1 uint64 `v8:"csReserved1"`
  CsReserved2 uint64 `v8:"csReserved2"`
}
