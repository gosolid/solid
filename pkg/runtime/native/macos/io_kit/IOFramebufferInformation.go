//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOFramebufferInformation
*/

type IOFramebufferInformation struct {
  BaseAddress uint64 `v8:"baseAddress"`
  ActiveWidth uint `v8:"activeWidth"`
  ActiveHeight uint `v8:"activeHeight"`
  BytesPerRow uint64 `v8:"bytesPerRow"`
  BytesPerPlane uint64 `v8:"bytesPerPlane"`
  BitsPerPixel uint `v8:"bitsPerPixel"`
  PixelType uint `v8:"pixelType"`
  Flags uint `v8:"flags"`
  Reserved [4]uint `v8:"reserved"`
}
