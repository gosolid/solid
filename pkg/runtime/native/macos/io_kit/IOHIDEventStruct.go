//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct IOKit.IOHIDEventStruct
*/

type IOHIDEventStruct struct {
  Type int `v8:"type"`
  ElementCookie uint `v8:"elementCookie"`
  Value int `v8:"value"`
  Timestamp objc.UnsignedWide `v8:"timestamp"`
  LongValueSize uint `v8:"longValueSize"`
  LongValue void `v8:"longValue"`
}
