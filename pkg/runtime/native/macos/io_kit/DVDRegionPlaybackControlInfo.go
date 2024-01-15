//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.DVDRegionPlaybackControlInfo
*/

type DVDRegionPlaybackControlInfo struct {
  DataLength [2]byte `v8:"dataLength"`
  Reserved [2]byte `v8:"reserved"`
  NumberUserResets byte `v8:"numberUserResets"`
  NumberVendorResets byte `v8:"numberVendorResets"`
  TypeCode byte `v8:"typeCode"`
  DriveRegion byte `v8:"driveRegion"`
  RpcScheme byte `v8:"rpcScheme"`
  Reserved2 byte `v8:"reserved2"`
}
