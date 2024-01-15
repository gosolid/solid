//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDDisplayTimingRangeRec
*/

type VDDisplayTimingRangeRec struct {
  CsRangeSize uint `v8:"csRangeSize"`
  CsRangeType uint `v8:"csRangeType"`
  CsRangeVersion uint `v8:"csRangeVersion"`
  CsRangeReserved uint `v8:"csRangeReserved"`
  CsRangeBlockIndex uint `v8:"csRangeBlockIndex"`
  CsRangeGroup uint `v8:"csRangeGroup"`
  CsRangeBlockCount uint `v8:"csRangeBlockCount"`
  CsRangeFlags uint `v8:"csRangeFlags"`
  CsMinPixelClock uint64 `v8:"csMinPixelClock"`
  CsMaxPixelClock uint64 `v8:"csMaxPixelClock"`
  CsMaxPixelError uint `v8:"csMaxPixelError"`
  CsTimingRangeSyncFlags uint `v8:"csTimingRangeSyncFlags"`
  CsTimingRangeSignalLevels uint `v8:"csTimingRangeSignalLevels"`
  CsTimingRangeSupportedSignalConfigs uint `v8:"csTimingRangeSupportedSignalConfigs"`
  CsMinFrameRate uint `v8:"csMinFrameRate"`
  CsMaxFrameRate uint `v8:"csMaxFrameRate"`
  CsMinLineRate uint `v8:"csMinLineRate"`
  CsMaxLineRate uint `v8:"csMaxLineRate"`
  CsMaxHorizontalTotal uint `v8:"csMaxHorizontalTotal"`
  CsMaxVerticalTotal uint `v8:"csMaxVerticalTotal"`
  CsMaxTotalReserved1 uint `v8:"csMaxTotalReserved1"`
  CsMaxTotalReserved2 uint `v8:"csMaxTotalReserved2"`
  CsCharSizeHorizontalActive byte `v8:"csCharSizeHorizontalActive"`
  CsCharSizeHorizontalBlanking byte `v8:"csCharSizeHorizontalBlanking"`
  CsCharSizeHorizontalSyncOffset byte `v8:"csCharSizeHorizontalSyncOffset"`
  CsCharSizeHorizontalSyncPulse byte `v8:"csCharSizeHorizontalSyncPulse"`
  CsCharSizeVerticalActive byte `v8:"csCharSizeVerticalActive"`
  CsCharSizeVerticalBlanking byte `v8:"csCharSizeVerticalBlanking"`
  CsCharSizeVerticalSyncOffset byte `v8:"csCharSizeVerticalSyncOffset"`
  CsCharSizeVerticalSyncPulse byte `v8:"csCharSizeVerticalSyncPulse"`
  CsCharSizeHorizontalBorderLeft byte `v8:"csCharSizeHorizontalBorderLeft"`
  CsCharSizeHorizontalBorderRight byte `v8:"csCharSizeHorizontalBorderRight"`
  CsCharSizeVerticalBorderTop byte `v8:"csCharSizeVerticalBorderTop"`
  CsCharSizeVerticalBorderBottom byte `v8:"csCharSizeVerticalBorderBottom"`
  CsCharSizeHorizontalTotal byte `v8:"csCharSizeHorizontalTotal"`
  CsCharSizeVerticalTotal byte `v8:"csCharSizeVerticalTotal"`
  CsCharSizeReserved1 uint16 `v8:"csCharSizeReserved1"`
  CsMinHorizontalActiveClocks uint `v8:"csMinHorizontalActiveClocks"`
  CsMaxHorizontalActiveClocks uint `v8:"csMaxHorizontalActiveClocks"`
  CsMinHorizontalBlankingClocks uint `v8:"csMinHorizontalBlankingClocks"`
  CsMaxHorizontalBlankingClocks uint `v8:"csMaxHorizontalBlankingClocks"`
  CsMinHorizontalSyncOffsetClocks uint `v8:"csMinHorizontalSyncOffsetClocks"`
  CsMaxHorizontalSyncOffsetClocks uint `v8:"csMaxHorizontalSyncOffsetClocks"`
  CsMinHorizontalPulseWidthClocks uint `v8:"csMinHorizontalPulseWidthClocks"`
  CsMaxHorizontalPulseWidthClocks uint `v8:"csMaxHorizontalPulseWidthClocks"`
  CsMinVerticalActiveClocks uint `v8:"csMinVerticalActiveClocks"`
  CsMaxVerticalActiveClocks uint `v8:"csMaxVerticalActiveClocks"`
  CsMinVerticalBlankingClocks uint `v8:"csMinVerticalBlankingClocks"`
  CsMaxVerticalBlankingClocks uint `v8:"csMaxVerticalBlankingClocks"`
  CsMinVerticalSyncOffsetClocks uint `v8:"csMinVerticalSyncOffsetClocks"`
  CsMaxVerticalSyncOffsetClocks uint `v8:"csMaxVerticalSyncOffsetClocks"`
  CsMinVerticalPulseWidthClocks uint `v8:"csMinVerticalPulseWidthClocks"`
  CsMaxVerticalPulseWidthClocks uint `v8:"csMaxVerticalPulseWidthClocks"`
  CsMinHorizontalBorderLeft uint `v8:"csMinHorizontalBorderLeft"`
  CsMaxHorizontalBorderLeft uint `v8:"csMaxHorizontalBorderLeft"`
  CsMinHorizontalBorderRight uint `v8:"csMinHorizontalBorderRight"`
  CsMaxHorizontalBorderRight uint `v8:"csMaxHorizontalBorderRight"`
  CsMinVerticalBorderTop uint `v8:"csMinVerticalBorderTop"`
  CsMaxVerticalBorderTop uint `v8:"csMaxVerticalBorderTop"`
  CsMinVerticalBorderBottom uint `v8:"csMinVerticalBorderBottom"`
  CsMaxVerticalBorderBottom uint `v8:"csMaxVerticalBorderBottom"`
  CsMaxNumLinks uint `v8:"csMaxNumLinks"`
  CsMinLink0PixelClock uint `v8:"csMinLink0PixelClock"`
  CsMaxLink0PixelClock uint `v8:"csMaxLink0PixelClock"`
  CsMinLink1PixelClock uint `v8:"csMinLink1PixelClock"`
  CsMaxLink1PixelClock uint `v8:"csMaxLink1PixelClock"`
  CsReserved6 uint `v8:"csReserved6"`
  CsReserved7 uint `v8:"csReserved7"`
  CsReserved8 uint `v8:"csReserved8"`
}
