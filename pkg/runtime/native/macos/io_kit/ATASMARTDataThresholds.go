//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.ATASMARTDataThresholds
*/

type ATASMARTDataThresholds struct {
  VendorSpecific1 [362]byte `v8:"vendorSpecific1"`
  VendorSpecific2 [149]byte `v8:"vendorSpecific2"`
  Checksum byte `v8:"checksum"`
}
