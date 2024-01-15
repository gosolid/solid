//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOFBDPLinkConfig
*/

type IOFBDPLinkConfig struct {
  Version uint16 `v8:"version"`
  BitRate byte `v8:"bitRate"`
  ReservedA [1]byte `v8:"reservedA"`
  T1Time uint16 `v8:"t1Time"`
  T2Time uint16 `v8:"t2Time"`
  T3Time uint16 `v8:"t3Time"`
  IdlePatterns byte `v8:"idlePatterns"`
  LaneCount byte `v8:"laneCount"`
  Voltage byte `v8:"voltage"`
  PreEmphasis byte `v8:"preEmphasis"`
  Downspread byte `v8:"downspread"`
  Scrambler byte `v8:"scrambler"`
  MaxBitRate byte `v8:"maxBitRate"`
  MaxLaneCount byte `v8:"maxLaneCount"`
  MaxDownspread byte `v8:"maxDownspread"`
  ReservedB [9]byte `v8:"reservedB"`
}
