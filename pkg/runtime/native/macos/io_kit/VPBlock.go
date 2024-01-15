//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct IOKit.VPBlock
*/

type VPBlock struct {
  VpBaseOffset uint `v8:"vpBaseOffset"`
  VpRowBytes uint `v8:"vpRowBytes"`
  VpBounds objc.Rect `v8:"vpBounds"`
  VpVersion int16 `v8:"vpVersion"`
  VpPackType int16 `v8:"vpPackType"`
  VpPackSize uint `v8:"vpPackSize"`
  VpHRes uint `v8:"vpHRes"`
  VpVRes uint `v8:"vpVRes"`
  VpPixelType int16 `v8:"vpPixelType"`
  VpPixelSize int16 `v8:"vpPixelSize"`
  VpCmpCount int16 `v8:"vpCmpCount"`
  VpCmpSize int16 `v8:"vpCmpSize"`
  VpPlaneBytes uint `v8:"vpPlaneBytes"`
}
