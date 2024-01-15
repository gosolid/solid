//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IODot3StatsEntry
*/

type IODot3StatsEntry struct {
  AlignmentErrors uint `v8:"alignmentErrors"`
  FcsErrors uint `v8:"fcsErrors"`
  SingleCollisionFrames uint `v8:"singleCollisionFrames"`
  MultipleCollisionFrames uint `v8:"multipleCollisionFrames"`
  SqeTestErrors uint `v8:"sqeTestErrors"`
  DeferredTransmissions uint `v8:"deferredTransmissions"`
  LateCollisions uint `v8:"lateCollisions"`
  ExcessiveCollisions uint `v8:"excessiveCollisions"`
  InternalMacTransmitErrors uint `v8:"internalMacTransmitErrors"`
  CarrierSenseErrors uint `v8:"carrierSenseErrors"`
  FrameTooLongs uint `v8:"frameTooLongs"`
  InternalMacReceiveErrors uint `v8:"internalMacReceiveErrors"`
  EtherChipSet uint `v8:"etherChipSet"`
  MissedFrames uint `v8:"missedFrames"`
}
