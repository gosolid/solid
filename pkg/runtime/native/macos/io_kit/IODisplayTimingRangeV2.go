//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IODisplayTimingRangeV2
*/

type IODisplayTimingRangeV2 struct {
  ReservedA [2]uint `v8:"reservedA"`
  Version uint `v8:"version"`
  ReservedB [5]uint `v8:"reservedB"`
  MinPixelClock uint64 `v8:"minPixelClock"`
  MaxPixelClock uint64 `v8:"maxPixelClock"`
  MaxPixelError uint `v8:"maxPixelError"`
  SupportedSyncFlags uint `v8:"supportedSyncFlags"`
  SupportedSignalLevels uint `v8:"supportedSignalLevels"`
  SupportedSignalConfigs uint `v8:"supportedSignalConfigs"`
  MinFrameRate uint `v8:"minFrameRate"`
  MaxFrameRate uint `v8:"maxFrameRate"`
  MinLineRate uint `v8:"minLineRate"`
  MaxLineRate uint `v8:"maxLineRate"`
  MaxHorizontalTotal uint `v8:"maxHorizontalTotal"`
  MaxVerticalTotal uint `v8:"maxVerticalTotal"`
  ReservedD [2]uint `v8:"reservedD"`
  CharSizeHorizontalActive byte `v8:"charSizeHorizontalActive"`
  CharSizeHorizontalBlanking byte `v8:"charSizeHorizontalBlanking"`
  CharSizeHorizontalSyncOffset byte `v8:"charSizeHorizontalSyncOffset"`
  CharSizeHorizontalSyncPulse byte `v8:"charSizeHorizontalSyncPulse"`
  CharSizeVerticalActive byte `v8:"charSizeVerticalActive"`
  CharSizeVerticalBlanking byte `v8:"charSizeVerticalBlanking"`
  CharSizeVerticalSyncOffset byte `v8:"charSizeVerticalSyncOffset"`
  CharSizeVerticalSyncPulse byte `v8:"charSizeVerticalSyncPulse"`
  CharSizeHorizontalBorderLeft byte `v8:"charSizeHorizontalBorderLeft"`
  CharSizeHorizontalBorderRight byte `v8:"charSizeHorizontalBorderRight"`
  CharSizeVerticalBorderTop byte `v8:"charSizeVerticalBorderTop"`
  CharSizeVerticalBorderBottom byte `v8:"charSizeVerticalBorderBottom"`
  CharSizeHorizontalTotal byte `v8:"charSizeHorizontalTotal"`
  CharSizeVerticalTotal byte `v8:"charSizeVerticalTotal"`
  ReservedE uint16 `v8:"reservedE"`
  MinHorizontalActiveClocks uint `v8:"minHorizontalActiveClocks"`
  MaxHorizontalActiveClocks uint `v8:"maxHorizontalActiveClocks"`
  MinHorizontalBlankingClocks uint `v8:"minHorizontalBlankingClocks"`
  MaxHorizontalBlankingClocks uint `v8:"maxHorizontalBlankingClocks"`
  MinHorizontalSyncOffsetClocks uint `v8:"minHorizontalSyncOffsetClocks"`
  MaxHorizontalSyncOffsetClocks uint `v8:"maxHorizontalSyncOffsetClocks"`
  MinHorizontalPulseWidthClocks uint `v8:"minHorizontalPulseWidthClocks"`
  MaxHorizontalPulseWidthClocks uint `v8:"maxHorizontalPulseWidthClocks"`
  MinVerticalActiveClocks uint `v8:"minVerticalActiveClocks"`
  MaxVerticalActiveClocks uint `v8:"maxVerticalActiveClocks"`
  MinVerticalBlankingClocks uint `v8:"minVerticalBlankingClocks"`
  MaxVerticalBlankingClocks uint `v8:"maxVerticalBlankingClocks"`
  MinVerticalSyncOffsetClocks uint `v8:"minVerticalSyncOffsetClocks"`
  MaxVerticalSyncOffsetClocks uint `v8:"maxVerticalSyncOffsetClocks"`
  MinVerticalPulseWidthClocks uint `v8:"minVerticalPulseWidthClocks"`
  MaxVerticalPulseWidthClocks uint `v8:"maxVerticalPulseWidthClocks"`
  MinHorizontalBorderLeft uint `v8:"minHorizontalBorderLeft"`
  MaxHorizontalBorderLeft uint `v8:"maxHorizontalBorderLeft"`
  MinHorizontalBorderRight uint `v8:"minHorizontalBorderRight"`
  MaxHorizontalBorderRight uint `v8:"maxHorizontalBorderRight"`
  MinVerticalBorderTop uint `v8:"minVerticalBorderTop"`
  MaxVerticalBorderTop uint `v8:"maxVerticalBorderTop"`
  MinVerticalBorderBottom uint `v8:"minVerticalBorderBottom"`
  MaxVerticalBorderBottom uint `v8:"maxVerticalBorderBottom"`
  MaxNumLinks uint `v8:"maxNumLinks"`
  MinLink0PixelClock uint `v8:"minLink0PixelClock"`
  MaxLink0PixelClock uint `v8:"maxLink0PixelClock"`
  MinLink1PixelClock uint `v8:"minLink1PixelClock"`
  MaxLink1PixelClock uint `v8:"maxLink1PixelClock"`
  SupportedPixelEncoding uint16 `v8:"supportedPixelEncoding"`
  SupportedBitsPerColorComponent uint16 `v8:"supportedBitsPerColorComponent"`
  SupportedColorimetryModes uint16 `v8:"supportedColorimetryModes"`
  SupportedDynamicRangeModes uint16 `v8:"supportedDynamicRangeModes"`
  ReservedF [1]uint `v8:"reservedF"`
  MaxBandwidth uint64 `v8:"maxBandwidth"`
  DscMinSliceHeight uint `v8:"dscMinSliceHeight"`
  DscMaxSliceHeight uint `v8:"dscMaxSliceHeight"`
  DscMinSliceWidth uint `v8:"dscMinSliceWidth"`
  DscMaxSliceWidth uint `v8:"dscMaxSliceWidth"`
  DscMinSlicePerLine uint `v8:"dscMinSlicePerLine"`
  DscMaxSlicePerLine uint `v8:"dscMaxSlicePerLine"`
  DscMinBPC uint16 `v8:"dscMinBPC"`
  DscMaxBPC uint16 `v8:"dscMaxBPC"`
  DscMinBPP uint16 `v8:"dscMinBPP"`
  DscMaxBPP uint16 `v8:"dscMaxBPP"`
  DscVBR byte `v8:"dscVBR"`
  DscBlockPredEnable byte `v8:"dscBlockPredEnable"`
  ReservedC [6]uint `v8:"reservedC"`
}
