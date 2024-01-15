//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDConvolutionInfoRec
*/

type VDConvolutionInfoRec struct {
  CsDisplayModeID int `v8:"csDisplayModeID"`
  CsDepthMode uint16 `v8:"csDepthMode"`
  CsPage uint `v8:"csPage"`
  CsFlags uint `v8:"csFlags"`
  CsReserved uint `v8:"csReserved"`
}
