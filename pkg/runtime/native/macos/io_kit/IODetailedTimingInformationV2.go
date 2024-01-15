//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IODetailedTimingInformationV2
*/

type IODetailedTimingInformationV2 struct {
  ReservedA [3]uint `v8:"reservedA"`
  HorizontalScaledInset uint `v8:"horizontalScaledInset"`
  VerticalScaledInset uint `v8:"verticalScaledInset"`
  ScalerFlags uint `v8:"scalerFlags"`
  HorizontalScaled uint `v8:"horizontalScaled"`
  VerticalScaled uint `v8:"verticalScaled"`
  SignalConfig uint `v8:"signalConfig"`
  SignalLevels uint `v8:"signalLevels"`
  PixelClock uint64 `v8:"pixelClock"`
  MinPixelClock uint64 `v8:"minPixelClock"`
  MaxPixelClock uint64 `v8:"maxPixelClock"`
  HorizontalActive uint `v8:"horizontalActive"`
  HorizontalBlanking uint `v8:"horizontalBlanking"`
  HorizontalSyncOffset uint `v8:"horizontalSyncOffset"`
  HorizontalSyncPulseWidth uint `v8:"horizontalSyncPulseWidth"`
  VerticalActive uint `v8:"verticalActive"`
  VerticalBlanking uint `v8:"verticalBlanking"`
  VerticalSyncOffset uint `v8:"verticalSyncOffset"`
  VerticalSyncPulseWidth uint `v8:"verticalSyncPulseWidth"`
  HorizontalBorderLeft uint `v8:"horizontalBorderLeft"`
  HorizontalBorderRight uint `v8:"horizontalBorderRight"`
  VerticalBorderTop uint `v8:"verticalBorderTop"`
  VerticalBorderBottom uint `v8:"verticalBorderBottom"`
  HorizontalSyncConfig uint `v8:"horizontalSyncConfig"`
  HorizontalSyncLevel uint `v8:"horizontalSyncLevel"`
  VerticalSyncConfig uint `v8:"verticalSyncConfig"`
  VerticalSyncLevel uint `v8:"verticalSyncLevel"`
  NumLinks uint `v8:"numLinks"`
  VerticalBlankingExtension uint `v8:"verticalBlankingExtension"`
  PixelEncoding uint16 `v8:"pixelEncoding"`
  BitsPerColorComponent uint16 `v8:"bitsPerColorComponent"`
  Colorimetry uint16 `v8:"colorimetry"`
  DynamicRange uint16 `v8:"dynamicRange"`
  DscCompressedBitsPerPixel uint16 `v8:"dscCompressedBitsPerPixel"`
  DscSliceHeight uint16 `v8:"dscSliceHeight"`
  DscSliceWidth uint16 `v8:"dscSliceWidth"`
  VerticalBlankingMaxStretchPerFrame uint16 `v8:"verticalBlankingMaxStretchPerFrame"`
  VerticalBlankingMaxShrinkPerFrame uint16 `v8:"verticalBlankingMaxShrinkPerFrame"`
  ReservedB [3]uint16 `v8:"reservedB"`
}
