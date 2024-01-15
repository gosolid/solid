//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDTimingInfoRec
*/

type VDTimingInfoRec struct {
  CsTimingMode int `v8:"csTimingMode"`
  CsTimingReserved uint64 `v8:"csTimingReserved"`
  CsTimingFormat uint `v8:"csTimingFormat"`
  CsTimingData uint `v8:"csTimingData"`
  CsTimingFlags uint `v8:"csTimingFlags"`
}
