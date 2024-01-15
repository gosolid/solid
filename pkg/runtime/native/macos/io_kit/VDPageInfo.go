//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDPageInfo
*/

type VDPageInfo struct {
  CsMode uint16 `v8:"csMode"`
  CsData int `v8:"csData"`
  CsPage int16 `v8:"csPage"`
  CsBaseAddr *byte `v8:"csBaseAddr"`
}
