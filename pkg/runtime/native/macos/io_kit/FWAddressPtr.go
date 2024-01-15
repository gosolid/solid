//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.FWAddressPtr
*/

type FWAddressPtr struct {
  NodeID uint16 `v8:"nodeID"`
  AddressHi uint16 `v8:"addressHi"`
  AddressLo uint `v8:"addressLo"`
}
