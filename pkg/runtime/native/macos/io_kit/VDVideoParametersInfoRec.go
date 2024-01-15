//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDVideoParametersInfoRec
*/

type VDVideoParametersInfoRec struct {
  CsDisplayModeID int `v8:"csDisplayModeID"`
  CsDepthMode uint16 `v8:"csDepthMode"`
  CsVPBlockPtr *VPBlock `v8:"csVPBlockPtr"`
  CsPageCount uint `v8:"csPageCount"`
  CsDeviceType uint `v8:"csDeviceType"`
  CsDepthFlags uint `v8:"csDepthFlags"`
}
