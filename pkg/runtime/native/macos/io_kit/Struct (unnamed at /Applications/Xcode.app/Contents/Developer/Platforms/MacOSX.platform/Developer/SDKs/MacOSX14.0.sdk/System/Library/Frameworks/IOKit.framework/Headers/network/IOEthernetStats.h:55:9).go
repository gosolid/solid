//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/IOKit.framework/Headers/network/IOEthernetStats.h:55:9)
*/

type Struct (unnamed at /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX14.0.sdk/System/Library/Frameworks/IOKit.framework/Headers/network/IOEthernetStats.h:55:9) struct {
  AlignmentErrors int `v8:"alignmentErrors"`
  FcsErrors int `v8:"fcsErrors"`
  SingleCollisionFrames int `v8:"singleCollisionFrames"`
  MultipleCollisionFrames int `v8:"multipleCollisionFrames"`
  SqeTestErrors int `v8:"sqeTestErrors"`
  DeferredTransmissions int `v8:"deferredTransmissions"`
  LateCollisions int `v8:"lateCollisions"`
  ExcessiveCollisions int `v8:"excessiveCollisions"`
  InternalMacTransmitErrors int `v8:"internalMacTransmitErrors"`
  CarrierSenseErrors int `v8:"carrierSenseErrors"`
  FrameTooLongs int `v8:"frameTooLongs"`
  InternalMacReceiveErrors int `v8:"internalMacReceiveErrors"`
  EtherChipSet int `v8:"etherChipSet"`
  MissedFrames int `v8:"missedFrames"`
}
