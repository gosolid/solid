//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.NXTabletProximityDataPtr
*/

type NXTabletProximityDataPtr struct {
  VendorID uint16 `v8:"vendorID"`
  TabletID uint16 `v8:"tabletID"`
  PointerID uint16 `v8:"pointerID"`
  DeviceID uint16 `v8:"deviceID"`
  SystemTabletID uint16 `v8:"systemTabletID"`
  VendorPointerType uint16 `v8:"vendorPointerType"`
  PointerSerialNumber uint `v8:"pointerSerialNumber"`
  UniqueID uint64 `v8:"uniqueID"`
  CapabilityMask uint `v8:"capabilityMask"`
  PointerType byte `v8:"pointerType"`
  EnterProximity byte `v8:"enterProximity"`
  Reserved1 int16 `v8:"reserved1"`
}
