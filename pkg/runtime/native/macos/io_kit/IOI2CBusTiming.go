//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct IOKit.IOI2CBusTiming
*/

type IOI2CBusTiming struct {
  BitTimeout objc.UnsignedWide `v8:"bitTimeout"`
  ByteTimeout objc.UnsignedWide `v8:"byteTimeout"`
  AcknowledgeTimeout objc.UnsignedWide `v8:"acknowledgeTimeout"`
  StartTimeout objc.UnsignedWide `v8:"startTimeout"`
  HoldTime objc.UnsignedWide `v8:"holdTime"`
  RiseFallTime objc.UnsignedWide `v8:"riseFallTime"`
  ReservedA [8]uint `v8:"reservedA"`
}
