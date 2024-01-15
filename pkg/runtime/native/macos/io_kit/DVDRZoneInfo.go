//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.DVDRZoneInfo
*/

type DVDRZoneInfo struct {
  DataLength uint16 `v8:"dataLength"`
  RzoneNumberLSB byte `v8:"rzoneNumberLSB"`
  BorderNumberLSB byte `v8:"borderNumberLSB"`
  Reserved byte `v8:"reserved"`
  Reserved2 byte `v8:"reserved2"`
  Copy byte `v8:"copy"`
  Damage byte `v8:"damage"`
  Reserved3 byte `v8:"reserved3"`
  Reserved4 byte `v8:"reserved4"`
  RestrictedOverwrite byte `v8:"restrictedOverwrite"`
  Incremental byte `v8:"incremental"`
  Blank byte `v8:"blank"`
  ReservedRZone byte `v8:"reservedRZone"`
  NextWritableAddressValid byte `v8:"nextWritableAddressValid"`
  LastRecordedAddressValid byte `v8:"lastRecordedAddressValid"`
  Reserved5 byte `v8:"reserved5"`
  RzoneStartAddress uint `v8:"rzoneStartAddress"`
  NextWritableAddress uint `v8:"nextWritableAddress"`
  FreeBlocks uint `v8:"freeBlocks"`
  BlockingFactor uint `v8:"blockingFactor"`
  RzoneSize uint `v8:"rzoneSize"`
  LastRecordedAddress uint `v8:"lastRecordedAddress"`
  RzoneNumberMSB byte `v8:"rzoneNumberMSB"`
  BorderNumberMSB byte `v8:"borderNumberMSB"`
  Reserved6 byte `v8:"reserved6"`
  Reserved7 byte `v8:"reserved7"`
}
