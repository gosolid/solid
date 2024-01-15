//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.DVDPhysicalFormatInfo
*/

type DVDPhysicalFormatInfo struct {
  DataLength [2]byte `v8:"dataLength"`
  Reserved [2]byte `v8:"reserved"`
  PartVersion byte `v8:"partVersion"`
  BookType byte `v8:"bookType"`
  MinimumRate byte `v8:"minimumRate"`
  DiscSize byte `v8:"discSize"`
  LayerType byte `v8:"layerType"`
  TrackPath byte `v8:"trackPath"`
  NumberOfLayers byte `v8:"numberOfLayers"`
  Reserved2 byte `v8:"reserved2"`
  TrackDensity byte `v8:"trackDensity"`
  LinearDensity byte `v8:"linearDensity"`
  Zero1 byte `v8:"zero1"`
  StartingPhysicalSectorNumberOfDataArea [3]byte `v8:"startingPhysicalSectorNumberOfDataArea"`
  Zero2 byte `v8:"zero2"`
  EndPhysicalSectorNumberOfDataArea [3]byte `v8:"endPhysicalSectorNumberOfDataArea"`
  Zero3 byte `v8:"zero3"`
  EndSectorNumberInLayerZero [3]byte `v8:"endSectorNumberInLayerZero"`
  Reserved1 byte `v8:"reserved1"`
  BcaFlag byte `v8:"bcaFlag"`
  MediaSpecific [2031]byte `v8:"mediaSpecific"`
}
