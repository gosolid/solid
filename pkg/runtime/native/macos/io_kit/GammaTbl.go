//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.GammaTbl
*/

type GammaTbl struct {
  GVersion int16 `v8:"gVersion"`
  GType int16 `v8:"gType"`
  GFormulaSize int16 `v8:"gFormulaSize"`
  GChanCnt int16 `v8:"gChanCnt"`
  GDataCnt int16 `v8:"gDataCnt"`
  GDataWidth int16 `v8:"gDataWidth"`
  GFormulaData [1]int16 `v8:"gFormulaData"`
}
