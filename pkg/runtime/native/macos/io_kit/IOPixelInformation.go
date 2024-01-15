//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOPixelInformation
*/

type IOPixelInformation struct {
  BytesPerRow uint `v8:"bytesPerRow"`
  BytesPerPlane uint `v8:"bytesPerPlane"`
  BitsPerPixel uint `v8:"bitsPerPixel"`
  PixelType uint `v8:"pixelType"`
  ComponentCount uint `v8:"componentCount"`
  BitsPerComponent uint `v8:"bitsPerComponent"`
  ComponentMasks [16]uint `v8:"componentMasks"`
  PixelFormat [64]byte `v8:"pixelFormat"`
  Flags uint `v8:"flags"`
  ActiveWidth uint `v8:"activeWidth"`
  ActiveHeight uint `v8:"activeHeight"`
  Reserved [2]uint `v8:"reserved"`
}
