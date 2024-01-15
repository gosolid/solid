//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDResolutionInfoRec
*/

type VDResolutionInfoRec struct {
  CsPreviousDisplayModeID int `v8:"csPreviousDisplayModeID"`
  CsDisplayModeID int `v8:"csDisplayModeID"`
  CsHorizontalPixels uint `v8:"csHorizontalPixels"`
  CsVerticalLines uint `v8:"csVerticalLines"`
  CsRefreshRate int `v8:"csRefreshRate"`
  CsMaxDepthMode uint16 `v8:"csMaxDepthMode"`
  CsResolutionFlags uint `v8:"csResolutionFlags"`
  CsReserved uint64 `v8:"csReserved"`
}
