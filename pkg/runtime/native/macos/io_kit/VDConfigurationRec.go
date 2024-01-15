//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDConfigurationRec
*/

type VDConfigurationRec struct {
  CsConfigFeature uint `v8:"csConfigFeature"`
  CsConfigSupport uint `v8:"csConfigSupport"`
  CsConfigValue uint64 `v8:"csConfigValue"`
  CsReserved1 uint64 `v8:"csReserved1"`
  CsReserved2 uint64 `v8:"csReserved2"`
}
