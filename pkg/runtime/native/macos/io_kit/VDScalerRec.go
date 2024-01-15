//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDScalerRec
*/

type VDScalerRec struct {
  CsScalerSize uint `v8:"csScalerSize"`
  CsScalerVersion uint `v8:"csScalerVersion"`
  CsReserved1 uint `v8:"csReserved1"`
  CsReserved2 uint `v8:"csReserved2"`
  CsDisplayModeID int `v8:"csDisplayModeID"`
  CsDisplayModeSeed uint `v8:"csDisplayModeSeed"`
  CsDisplayModeState uint `v8:"csDisplayModeState"`
  CsReserved3 uint `v8:"csReserved3"`
  CsScalerFlags uint `v8:"csScalerFlags"`
  CsHorizontalPixels uint `v8:"csHorizontalPixels"`
  CsVerticalPixels uint `v8:"csVerticalPixels"`
  CsHorizontalInset uint `v8:"csHorizontalInset"`
  CsVerticalInset uint `v8:"csVerticalInset"`
  CsReserved6 uint `v8:"csReserved6"`
  CsReserved7 uint `v8:"csReserved7"`
  CsReserved8 uint `v8:"csReserved8"`
}
