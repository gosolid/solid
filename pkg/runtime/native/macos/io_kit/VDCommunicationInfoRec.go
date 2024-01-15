//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDCommunicationInfoRec
*/

type VDCommunicationInfoRec struct {
  CsBusID int `v8:"csBusID"`
  CsBusType uint `v8:"csBusType"`
  CsMinBus int `v8:"csMinBus"`
  CsMaxBus int `v8:"csMaxBus"`
  CsSupportedTypes uint `v8:"csSupportedTypes"`
  CsSupportedCommFlags uint `v8:"csSupportedCommFlags"`
  CsReserved2 uint `v8:"csReserved2"`
  CsReserved3 uint `v8:"csReserved3"`
  CsReserved4 uint `v8:"csReserved4"`
  CsReserved5 uint `v8:"csReserved5"`
  CsReserved6 uint `v8:"csReserved6"`
  CsReserved7 uint `v8:"csReserved7"`
}
