//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.CDTrackInfo
*/

type CDTrackInfo struct {
  DataLength uint16 `v8:"dataLength"`
  TrackNumberLSB byte `v8:"trackNumberLSB"`
  SessionNumberLSB byte `v8:"sessionNumberLSB"`
  Reserved byte `v8:"reserved"`
  TrackMode byte `v8:"trackMode"`
  Copy byte `v8:"copy"`
  Damage byte `v8:"damage"`
  Reserved3 byte `v8:"reserved3"`
  DataMode byte `v8:"dataMode"`
  FixedPacket byte `v8:"fixedPacket"`
  Packet byte `v8:"packet"`
  Blank byte `v8:"blank"`
  ReservedTrack byte `v8:"reservedTrack"`
  NextWritableAddressValid byte `v8:"nextWritableAddressValid"`
  LastRecordedAddressValid byte `v8:"lastRecordedAddressValid"`
  Reserved5 byte `v8:"reserved5"`
  TrackStartAddress uint `v8:"trackStartAddress"`
  NextWritableAddress uint `v8:"nextWritableAddress"`
  FreeBlocks uint `v8:"freeBlocks"`
  FixedPacketSize uint `v8:"fixedPacketSize"`
  TrackSize uint `v8:"trackSize"`
  LastRecordedAddress uint `v8:"lastRecordedAddress"`
  TrackNumberMSB byte `v8:"trackNumberMSB"`
  SessionNumberMSB byte `v8:"sessionNumberMSB"`
  Reserved6 byte `v8:"reserved6"`
  Reserved7 byte `v8:"reserved7"`
}
