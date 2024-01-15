//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOBlitSurface
*/

type IOBlitSurface struct {
  Memory void `v8:"memory"`
  PixelFormat uint `v8:"pixelFormat"`
  Size IOBlitRectangle `v8:"size"`
  RowBytes uint `v8:"rowBytes"`
  ByteOffset uint `v8:"byteOffset"`
  Palette uint `v8:"palette"`
  AccessFlags uint `v8:"accessFlags"`
  InterfaceRef *IOBlitMemory `v8:"interfaceRef"`
  More [14]uint `v8:"more"`
}
