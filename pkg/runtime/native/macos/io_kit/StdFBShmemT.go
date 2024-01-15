//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct IOKit.StdFBShmem_t
*/

type StdFBShmemT struct {
  CursorSema int `v8:"cursorSema"`
  Frame int `v8:"frame"`
  CursorShow byte `v8:"cursorShow"`
  CursorObscured byte `v8:"cursorObscured"`
  ShieldFlag byte `v8:"shieldFlag"`
  Shielded byte `v8:"shielded"`
  SaveRect IOGBounds `v8:"saveRect"`
  ShieldRect IOGBounds `v8:"shieldRect"`
  CursorLoc IOGPoint `v8:"cursorLoc"`
  CursorRect IOGBounds `v8:"cursorRect"`
  OldCursorRect IOGBounds `v8:"oldCursorRect"`
  ScreenBounds IOGBounds `v8:"screenBounds"`
  Version int `v8:"version"`
  StructSize int `v8:"structSize"`
  VblTime objc.UnsignedWide `v8:"vblTime"`
  VblDelta objc.UnsignedWide `v8:"vblDelta"`
  VblCount uint64 `v8:"vblCount"`
  ReservedC [27]uint `v8:"reservedC"`
  HardwareCursorFlags [4]byte `v8:"hardwareCursorFlags"`
  HardwareCursorCapable byte `v8:"hardwareCursorCapable"`
  HardwareCursorActive byte `v8:"hardwareCursorActive"`
  HardwareCursorShields byte `v8:"hardwareCursorShields"`
  ReservedB [1]byte `v8:"reservedB"`
  CursorSize [4]IOGSize `v8:"cursorSize"`
  HotSpot [4]IOGPoint `v8:"hotSpot"`
  Cursor void `v8:"cursor"`
}
