//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDSizeInfo
*/

type VDSizeInfo struct {
  CsHSize int16 `v8:"csHSize"`
  CsHPos int16 `v8:"csHPos"`
  CsVSize int16 `v8:"csVSize"`
  CsVPos int16 `v8:"csVPos"`
}
