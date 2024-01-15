//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDScalerInfoRec
*/

type VDScalerInfoRec struct {
  CsScalerInfoSize uint `v8:"csScalerInfoSize"`
  CsScalerInfoVersion uint `v8:"csScalerInfoVersion"`
  CsReserved1 uint `v8:"csReserved1"`
  CsReserved2 uint `v8:"csReserved2"`
  CsScalerFeatures uint `v8:"csScalerFeatures"`
  CsMaxHorizontalPixels uint `v8:"csMaxHorizontalPixels"`
  CsMaxVerticalPixels uint `v8:"csMaxVerticalPixels"`
  CsReserved3 uint `v8:"csReserved3"`
  CsReserved4 uint `v8:"csReserved4"`
  CsReserved5 uint `v8:"csReserved5"`
  CsReserved6 uint `v8:"csReserved6"`
  CsReserved7 uint `v8:"csReserved7"`
}
