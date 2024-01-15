//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDDetailedTimingRec
*/

type VDDetailedTimingRec struct {
  CsTimingSize uint `v8:"csTimingSize"`
  CsTimingType uint `v8:"csTimingType"`
  CsTimingVersion uint `v8:"csTimingVersion"`
  CsTimingReserved uint `v8:"csTimingReserved"`
  CsDisplayModeID int `v8:"csDisplayModeID"`
  CsDisplayModeSeed uint `v8:"csDisplayModeSeed"`
  CsDisplayModeState uint `v8:"csDisplayModeState"`
  CsDisplayModeAlias uint `v8:"csDisplayModeAlias"`
  CsSignalConfig uint `v8:"csSignalConfig"`
  CsSignalLevels uint `v8:"csSignalLevels"`
  CsPixelClock uint64 `v8:"csPixelClock"`
  CsMinPixelClock uint64 `v8:"csMinPixelClock"`
  CsMaxPixelClock uint64 `v8:"csMaxPixelClock"`
  CsHorizontalActive uint `v8:"csHorizontalActive"`
  CsHorizontalBlanking uint `v8:"csHorizontalBlanking"`
  CsHorizontalSyncOffset uint `v8:"csHorizontalSyncOffset"`
  CsHorizontalSyncPulseWidth uint `v8:"csHorizontalSyncPulseWidth"`
  CsVerticalActive uint `v8:"csVerticalActive"`
  CsVerticalBlanking uint `v8:"csVerticalBlanking"`
  CsVerticalSyncOffset uint `v8:"csVerticalSyncOffset"`
  CsVerticalSyncPulseWidth uint `v8:"csVerticalSyncPulseWidth"`
  CsHorizontalBorderLeft uint `v8:"csHorizontalBorderLeft"`
  CsHorizontalBorderRight uint `v8:"csHorizontalBorderRight"`
  CsVerticalBorderTop uint `v8:"csVerticalBorderTop"`
  CsVerticalBorderBottom uint `v8:"csVerticalBorderBottom"`
  CsHorizontalSyncConfig uint `v8:"csHorizontalSyncConfig"`
  CsHorizontalSyncLevel uint `v8:"csHorizontalSyncLevel"`
  CsVerticalSyncConfig uint `v8:"csVerticalSyncConfig"`
  CsVerticalSyncLevel uint `v8:"csVerticalSyncLevel"`
  CsNumLinks uint `v8:"csNumLinks"`
  CsReserved2 uint `v8:"csReserved2"`
  CsReserved3 uint `v8:"csReserved3"`
  CsReserved4 uint `v8:"csReserved4"`
  CsReserved5 uint `v8:"csReserved5"`
  CsReserved6 uint `v8:"csReserved6"`
  CsReserved7 uint `v8:"csReserved7"`
  CsReserved8 uint `v8:"csReserved8"`
}
